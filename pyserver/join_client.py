# pyserver/join_client.py
import os
import grpc
import swim_pb2
import swim_pb2_grpc

def join_system(new_node_id, new_node_address, bootstrap_address, bootstrap_node_id):
    print(f"Component NewNode of Node {new_node_id} sends RPC Join to Component Dissemination of Node {bootstrap_node_id}")
    with grpc.insecure_channel(bootstrap_address) as channel:
        stub = swim_pb2_grpc.DisseminationStub(channel)
        join_req = swim_pb2.JoinRequest(
            new_node_id=new_node_id,
            new_node_address=new_node_address
        )
        try:
            join_resp = stub.Join(join_req)
            return join_resp.membership  # a list of membership strings
        except grpc.RpcError as e:
            print(f"Error calling Join RPC: {e}")
            return []

if __name__ == "__main__":
    new_node_id = os.environ.get("NEW_NODE_ID", "6")
    new_node_address = os.environ.get("NEW_NODE_ADDRESS", "node6:50056")
    bootstrap_address = os.environ.get("BOOTSTRAP_ADDRESS", "localhost:50061")
    bootstrap_node_id = os.environ.get("BOOTSTRAP_NODE_ID", "1")
    membership = join_system(new_node_id, new_node_address, bootstrap_address, bootstrap_node_id)
    # Print as comma-separated list so that start.sh can capture it.
    print(",".join(membership))
