package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"google.golang.org/grpc"

	pb "failure-detector/proto"
)

// DisseminatorServer implements the Dissemination service.
type DisseminatorServer struct {
	pb.UnimplementedDisseminationServer
	nodeID     string
	membership []string // should contain failure detector addresses, e.g., "node1:50051"
}

// updateLocalFailureDetector updates the local Failure Detector's membership list.
func updateLocalFailureDetector(failedNode, nodeID string) error {
	// Calculate the failure detector's port: 50050 + nodeID.
	intNodeID, err := strconv.Atoi(nodeID)
	if err != nil {
		return err
	}
	fdPort := 50050 + intNodeID
	addr := fmt.Sprintf("localhost:%d", fdPort)

	// Dial the local Failure Detector service.
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to dial local failure detector: %v", err)
	}
	defer conn.Close()

	// Create a client for the FailureDetector service.
	fdClient := pb.NewFailureDetectorClient(conn)
	req := &pb.MembershipUpdateRequest{FailedNodeId: failedNode}
	resp, err := fdClient.UpdateMembership(context.Background(), req)
	if err != nil {
		return fmt.Errorf("failed to update membership: %v", err)
	}
	if !resp.Success {
		return errors.New("update membership RPC returned failure")
	}
	log.Printf("Local Failure Detector updated membership, removed %s", failedNode)
	return nil
}

// updateRemoteFailureDetector updates a remote node's Failure Detector service.
func updateRemoteFailureDetector(remoteAddr, failedNode string) error {
	// remoteAddr should be something like "node2:50052" (the address of a remote Failure Detector)
	conn, err := grpc.Dial(remoteAddr, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to dial remote failure detector at %s: %v", remoteAddr, err)
	}
	defer conn.Close()

	client := pb.NewFailureDetectorClient(conn)
	req := &pb.MembershipUpdateRequest{FailedNodeId: failedNode}
	resp, err := client.UpdateMembership(context.Background(), req)
	if err != nil {
		return fmt.Errorf("failed to update remote membership at %s: %v", remoteAddr, err)
	}
	if !resp.Success {
		return fmt.Errorf("remote update at %s returned failure", remoteAddr)
	}
	return nil
}

// Disseminate handles failure notifications.
func (s *DisseminatorServer) Disseminate(ctx context.Context, req *pb.DisseminationRequest) (*pb.DisseminationResponse, error) {
	// Server-side logging.
	log.Printf("Component Dissemination of Node %s runs RPC Disseminate called by Component FailureDetector of Node %s", s.nodeID, req.SenderId)
	fmt.Printf("Node %s detected failure of node %s\n", s.nodeID, req.FailedNodeId)

	// Update the local Failure Detector's membership list.
	if err := updateLocalFailureDetector(req.FailedNodeId, s.nodeID); err != nil {
		log.Printf("Failed to update local failure detector: %v", err)
	}

	// Propagate the update to remote nodes.
	for _, member := range s.membership {
		// Skip self. Assuming member addresses are formatted like "nodeX:50051".
		if strings.HasPrefix(member, fmt.Sprintf("node%s:", s.nodeID)) {
			continue
		}
		err := updateRemoteFailureDetector(member, req.FailedNodeId)
		if err != nil {
			log.Printf("Failed to update remote failure detector at %s: %v", member, err)
		} else {
			log.Printf("Successfully updated remote failure detector at %s; removed %s", member, req.FailedNodeId)
		}
	}

	return &pb.DisseminationResponse{Success: true}, nil
}

// Join handles join requests from new nodes.
func (s *DisseminatorServer) Join(ctx context.Context, req *pb.JoinRequest) (*pb.JoinResponse, error) {
	// Server-side logging.
	log.Printf("Component Dissemination of Node %s runs RPC Join called by Component NewNode of Node %s", s.nodeID, req.NewNodeId)
	// Respond with the current membership list.
	return &pb.JoinResponse{Membership: s.membership}, nil
}

func main() {
	// Read configuration from environment variables.
	nodeID := os.Getenv("NODE_ID")
	if nodeID == "" {
		log.Fatal("NODE_ID not set")
	}
	// Set membership to the failure detector addresses.
	// For example: "node1:50051", "node2:50052", ...
	membership := []string{"node1:50051", "node2:50052", "node3:50053", "node4:50054", "node5:50055"}

	// Compute the port for the Dissemination service.
	// Dissemination service listens on 50060 + nodeID.
	intNodeID, err := strconv.Atoi(nodeID)
	if err != nil {
		log.Fatalf("Invalid NODE_ID: %v", err)
	}
	portNum := 50060 + intNodeID
	address := fmt.Sprintf(":%d", portNum)

	// Create a TCP listener.
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", address, err)
	}

	// Create the gRPC server.
	grpcServer := grpc.NewServer()
	server := &DisseminatorServer{
		nodeID:     nodeID,
		membership: membership,
	}
	pb.RegisterDisseminationServer(grpcServer, server)

	log.Printf("Component Dissemination of Node %s listening on port %d", nodeID, portNum)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
