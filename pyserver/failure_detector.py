import grpc
import time
import random
from concurrent import futures
import pyserver.swim_pb2 as swim_pb2
import pyserver.swim_pb2_grpc as swim_pb2_grpc

T_PRIME = 5  # Interval for pinging in seconds
K = 3  # Number of indirect probe nodes

class FailureDetector(swim_pb2_grpc.FailureDetectorServicer):
    def __init__(self, node_id, members):
        self.node_id = node_id
        self.members = members  # List of other nodes' addresses
        self.last_heard = {node: time.time() for node in members}

    def Ping(self, request, context):
        print(f"Component FailureDetector of Node {self.node_id} runs RPC Ping called by Component FailureDetector of Node {request.sender_id}")
        return swim_pb2.PingResponse(is_alive=True)

    def IndirectPing(self, request, context):
        print(f"Component FailureDetector of Node {self.node_id} runs RPC IndirectPing called by Component FailureDetector of Node {request.requester_id}")
        return swim_pb2.IndirectPingResponse(success=True)

    def monitor_nodes(self):
        while True:
            time.sleep(T_PRIME)
            target = random.choice(self.members)
            print(f"Node {self.node_id} pinging {target}...")
            with grpc.insecure_channel(target) as channel:
                stub = swim_pb2_grpc.FailureDetectorStub(channel)
                try:
                    response = stub.Ping(swim_pb2.PingRequest(sender_id=self.node_id, target_id=target))
                    if response.is_alive:
                        self.last_heard[target] = time.time()
                except grpc.RpcError:
                    print(f"Node {target} did not respond, trying indirect probe...")
                    proxies = random.sample([m for m in self.members if m != target], min(K, len(self.members)-1))
                    for proxy in proxies:
                        try:
                            with grpc.insecure_channel(proxy) as pchannel:
                                pstub = swim_pb2_grpc.FailureDetectorStub(pchannel)
                                presp = pstub.IndirectPing(swim_pb2.IndirectPingRequest(requester_id=self.node_id, target_id=target, proxy_nodes=proxies))
                                if presp.success:
                                    self.last_heard[target] = time.time()
                                    break
                        except grpc.RpcError:
                            pass
                    else:
                        print(f"Node {target} is marked as failed!")

def serve(node_id, members):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    detector = FailureDetector(node_id, members)
    swim_pb2_grpc.add_FailureDetectorServicer_to_server(detector, server)
    server.add_insecure_port(f"[::]:{50050 + int(node_id)}")
    server.start()
    detector.monitor_nodes()
    server.wait_for_termination()

if __name__ == "__main__":
    node_id = input("Enter node ID: ")
    members = ["localhost:50051", "localhost:50052", "localhost:50053"]  # Replace with real addresses
    serve(node_id, members)