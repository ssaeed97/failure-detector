import grpc
import time
import random
import os
from concurrent import futures
import swim_pb2 as swim_pb2
import swim_pb2_grpc as swim_pb2_grpc

T_PRIME = 10 # Interval for pinging in seconds
K = 3        # Number of indirect probe nodes

def get_node_id_from_address(address):
    # Assumes port = 50050 + node_id, e.g., node1:50051 -> node_id = 1
    port = int(address.split(":")[-1])
    return str(port - 50050)

class FailureDetector(swim_pb2_grpc.FailureDetectorServicer):
    def __init__(self, node_id, members):
        self.node_id = node_id
        self.members = members  # List of nodes' addresses
        self.last_heard = {node: time.time() for node in members}

    def update_membership(self, failed_node):
        if failed_node in self.members:
            self.members.remove(failed_node)
            print(f"Membership updated: removed failed node {failed_node}. New membership: {self.members}")
    
    def Ping(self, request, context):
        # Server-side logging for Ping RPC
        print(f"Component FailureDetector of Node {self.node_id} runs RPC Ping called by Component FailureDetector of Node {request.sender_id}")
        return swim_pb2.PingResponse(is_alive=True)

    def IndirectPing(self, request, context):
        # Server-side logging for IndirectPing RPC
        print(f"Component FailureDetector of Node {self.node_id} runs RPC IndirectPing called by Component FailureDetector of Node {request.requester_id}")
        
        target_address = request.target_id
        try:
            # Create a channel to the target node.
            with grpc.insecure_channel(target_address) as channel:
                stub = swim_pb2_grpc.FailureDetectorStub(channel)
                target_node_id = get_node_id_from_address(target_address)
                print(f"Component FailureDetector of Node {self.node_id} sends RPC Ping to Component FailureDetector of Node {target_node_id}")
                
                # Send a Ping RPC to the target node.
                response = stub.Ping(swim_pb2.PingRequest(sender_id=self.node_id, target_id=target_address))
                if response.is_alive:
                    return swim_pb2.IndirectPingResponse(success=True)
                else:
                    return swim_pb2.IndirectPingResponse(success=False)
        except grpc.RpcError as e:
            # Extract error details
            #error_code = e.code()
            #error_details = e.details()
            print(f"IndirectPing: Failed to contact target {target_address} from proxy Node {self.node_id}: ")#Code: {error_code}, Details: {error_details}")
            return swim_pb2.IndirectPingResponse(success=False)



    def notify_failure(self, target):
        # Iterate over each member in the membership list.
        for member in self.members:
            # Skip self if the member address matches our own.
            if member == f"localhost:{50050 + int(self.node_id)}":
                continue
            # Extract the node ID from the member address.
            member_node_id = get_node_id_from_address(member)
            # Compute the dissemination port for that member (e.g., port = 50060 + member_node_id).
            dissemination_port = 50060 + int(member_node_id)
            dissemination_address = f"localhost:{dissemination_port}"
            # Client-side logging for the dissemination call.
            print(f"Component FailureDetector of Node {self.node_id} sends RPC Disseminate to Component Dissemination of Node {member_node_id}")
            with grpc.insecure_channel(dissemination_address) as channel:
                dissemination_stub = swim_pb2_grpc.DisseminationStub(channel)
                try:
                    response = dissemination_stub.Disseminate(
                        swim_pb2.DisseminationRequest(
                            sender_id=self.node_id,
                            failed_node_id=target
                        )
                    )
                    if response.success:
                        print(f"Dissemination to Node {member_node_id} successful for node {target}")
                    else:
                        print(f"Dissemination to Node {member_node_id} failed for node {target}")
                except grpc.RpcError as e:
                    print(f"Error calling Dissemination service on node {member_node_id}: {e}")
        # Update local membership list after dissemination.
        self.update_membership(target)

    def monitor_nodes(self):
        while True:
            time.sleep(T_PRIME)
            target = random.choice(self.members)
            target_node_id = get_node_id_from_address(target)
            print(f"\nNode {self.node_id} pinging {target}...")
            with grpc.insecure_channel(target) as channel:
                stub = swim_pb2_grpc.FailureDetectorStub(channel)
                try:
                    # Client-side logging before Ping RPC call
                    print(f"Component FailureDetector of Node {self.node_id} sends RPC Ping to Component FailureDetector of Node {target_node_id}")
                    response = stub.Ping(swim_pb2.PingRequest(sender_id=self.node_id, target_id=target))
                    if response.is_alive:
                        print("Node Alive, checked with Direct Ping")
                        self.last_heard[target] = time.time()
                except grpc.RpcError:
                    print(f"Node {target} did not respond, trying indirect probe...")
                    available_proxies = [m for m in self.members if m != target]
                    #print("AVAILABLE PROXIES")
                    #print(available_proxies)
                    proxies = random.sample(available_proxies, min(K, len(available_proxies)))
                    #print("PROXIES")
                    #print(proxies)
                    for proxy in proxies:
                        proxy_node_id = get_node_id_from_address(proxy)
                        # Client-side logging before IndirectPing RPC call
                        print(f"Component FailureDetector of Node {self.node_id} sends RPC IndirectPing to Component FailureDetector of Node {proxy_node_id}")
                        try:
                            with grpc.insecure_channel(proxy) as pchannel:
                                pstub = swim_pb2_grpc.FailureDetectorStub(pchannel)
                                presp = pstub.IndirectPing(
                                    swim_pb2.IndirectPingRequest(
                                        requester_id=self.node_id,
                                        target_id=target,
                                        proxy_nodes=proxies
                                    )
                                )
                                if presp.success:
                                    self.last_heard[target] = time.time()
                                    print("Indirect Ping was success, Node Alive")
                                    break
                        except grpc.RpcError:
                            continue
                    else:
                        print(f"Node {target} is marked as failed!")
                        # Call the dissemination service to multicast the failure
                        self.notify_failure(target)

def serve(node_id, members):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    detector = FailureDetector(node_id, members)
    swim_pb2_grpc.add_FailureDetectorServicer_to_server(detector, server)
    # Bind to port = 50050 + node_id
    server.add_insecure_port(f"[::]:{50050 + int(node_id)}")
    server.start()
    detector.monitor_nodes()
    server.wait_for_termination()

if __name__ == "__main__":
    # For containerized environments, read from environment variables.
    node_id = os.environ.get("NODE_ID")
    if not node_id:
        node_id = input("Enter node ID: ")

    members_env = os.environ.get("MEMBERS")
    if members_env:
        members = members_env.split(',')
    else:
        members = ["localhost:50051", "localhost:50052", "localhost:50053", "localhost:50054", "localhost:50055"]
    
    serve(node_id, members)
