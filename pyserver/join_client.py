import grpc
import os
import swim_pb2
import swim_pb2_grpc

def join_system(new_node_id, new_node_address, bootstrap_address, bootstrap_node_id):
    # Client-side logging before calling the Join RPC.
    print(f"Component NewNode of Node {new_node_id} sends RPC Join to Component Dissemination of Node {bootstrap_node_id}")
    
    # Establish a gRPC channel to the bootstrap node's Dissemination service.
    with grpc.insecure_channel(bootstrap_address) as channel:
        dissemination_stub = swim_pb2_grpc.DisseminationStub(channel)
        # Build the JoinRequest message.
        join_req = swim_pb2.JoinRequest(
            new_node_id=new_node_id,
            new_node_address=new_node_address
        )
        try:
            # Make the Join RPC call.
            join_resp = dissemination_stub.Join(join_req)
            print(f"Received membership list: {join_resp.membership}")
        except grpc.RpcError as e:
            print(f"Error calling Join RPC: {e}")

if __name__ == "__main__":
    # These values could be set via environment variables or command-line arguments.
    # For example, assume:
    # - The new node has ID 4, address "node4:50064" (since port = 50060 + node_id).
    # - The bootstrap node is node1, listening on "node1:50061".
    # - The bootstrap node's ID is "1".

    new_node_id = os.environ.get("NEW_NODE_ID", "4")
    new_node_address = os.environ.get("NEW_NODE_ADDRESS", "node4:50064")
    bootstrap_address = os.environ.get("BOOTSTRAP_ADDRESS", "localhost:50061")
    bootstrap_node_id = os.environ.get("BOOTSTRAP_NODE_ID", "1")
    
    join_system(new_node_id, new_node_address, bootstrap_address, bootstrap_node_id)
