import grpc
import time
import random
import os
from concurrent import futures
import swim_pb2 as swim_pb2
import swim_pb2_grpc as swim_pb2_grpc

T_PRIME = 5  # Interval for pinging in seconds
K = 3        # Number of indirect probe nodes

def get_node_id_from_address(address):
    # Assumes port = 50050 + node_id, e.g., localhost:50051 -> node_id = 1
    port = int(address.split(":")[-1])
    return str(port - 50050)

class FailureDetector(swim_pb2_grpc.FailureDetectorServicer):
    def __init__(self, node_id, members):
        self.node_id = node_id
        self.members = members  # List of nodes' addresses
        self.last_heard = {node: time.time() for node in members}

    def Ping(self, request, context):
        # Server-side logging for Ping RPC
        print(f"Component FailureDetector of Node {self.node_id} runs RPC Ping called by Component FailureDetector of Node {request.sender_id}")
        return swim_pb2.PingResponse(is_alive=True)

    def IndirectPing(self, request, context):
        # Server-side logging for IndirectPing RPC
        print(f"Component FailureDetector of Node {self.node_id} runs RPC IndirectPing called by Component FailureDetector of Node {request.requester_id}")
        return swim_pb2.IndirectPingResponse(success=True)

    def monitor_nodes(self):
        while True:
            time.sleep(T_PRIME)
            target = random.choice(self.members)
            target_node_id = get_node_id_from_address(target)
            print(f"Node {self.node_id} pinging {target}...")
            with grpc.insecure_channel(target) as channel:
                stub = swim_pb2_grpc.FailureDetectorStub(channel)
                try:
                    # Client-side logging before Ping RPC call
                    print(f"Component FailureDetector of Node {self.node_id} sends RPC Ping to Component FailureDetector of Node {target_node_id}")
                    response = stub.Ping(swim_pb2.PingRequest(sender_id=self.node_id, target_id=target))
                    if response.is_alive:
                        self.last_heard[target] = time.time()
                except grpc.RpcError:
                    print(f"Node {target} did not respond, trying indirect probe...")
                    available_proxies = [m for m in self.members if m != target]
                    proxies = random.sample(available_proxies, min(K, len(available_proxies)))
                    for proxy in proxies:
                        proxy_node_id = get_node_id_from_address(proxy)
                        # Client-side logging before IndirectPing RPC call
                        print(f"Component FailureDetector of Node {self.node_id} sends RPC IndirectPing to Component FailureDetector of Node {proxy_node_id}")
                        try:
                            with grpc.insecure_channel(proxy) as pchannel:
                                pstub = swim_pb2_grpc.FailureDetectorStub(pchannel)
                                presp = pstub.IndirectPing(swim_pb2.IndirectPingRequest(
                                    requester_id=self.node_id,
                                    target_id=target,
                                    proxy_nodes=proxies))
                                if presp.success:
                                    self.last_heard[target] = time.time()
                                    break
                        except grpc.RpcError:
                            continue
                    else:
                        print(f"Node {target} is marked as failed!")

def serve(node_id, members):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    detector = FailureDetector(node_id, members)
    swim_pb2_grpc.add_FailureDetectorServicer_to_server(detector, server)
    # Bind to port = 50050 + node_id
    server.add_insecure_port(f"[::]:{50050 + int(node_id)}")
    server.start()
    detector.monitor_nodes()
    server.wait_for_termination()
#Code for Docker
if __name__ == "__main__":
    # Read NODE_ID from environment variable if available; otherwise, prompt for input.
    node_id = os.environ.get("NODE_ID")
    if not node_id:
        node_id = input("Enter node ID: ")

    # Read MEMBERS from environment variable; if not provided, default to hardcoded values.
    members_env = os.environ.get("MEMBERS")
    if members_env:
        members = members_env.split(',')
    else:
        members = ["node1:50051", "node2:50052", "node3:50053"]
    
    serve(node_id, members)

#Code for running locally
'''if __name__ == "__main__":
    node_id = input("Enter node ID: ")
    members = ["localhost:50051", "localhost:50052", "localhost:50053"]  # Replace with actual addresses as needed
    serve(node_id, members)'''
