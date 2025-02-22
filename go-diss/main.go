package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"google.golang.org/grpc"

	// Import the generated proto package.
	// Adjust the import path based on the go_package option in your proto file.
	pb "go-diss/proto"
)

// DisseminatorServer implements the Dissemination service.
type DisseminatorServer struct {
	pb.UnimplementedDisseminationServer
	nodeID     string
	membership []string
}

// Disseminate handles failure notifications.
func (s *DisseminatorServer) Disseminate(ctx context.Context, req *pb.DisseminationRequest) (*pb.DisseminationResponse, error) {
	// Server-side logging
	log.Printf("\nComponent Dissemination of Node %s runs RPC Disseminate called by Component FailureDetector of Node %s", s.nodeID, req.SenderId)
	// Process the failed node notification (e.g., update internal state)
	fmt.Printf("Node %s detected failure of node %s\n", s.nodeID, req.FailedNodeId)
	return &pb.DisseminationResponse{Success: true}, nil
}

// Join handles join requests from new nodes.
func (s *DisseminatorServer) Join(ctx context.Context, req *pb.JoinRequest) (*pb.JoinResponse, error) {
	// Server-side logging
	log.Printf("Component Dissemination of Node %s runs RPC Join called by Component NewNode of Node %s", s.nodeID, req.NewNodeId)
	// Respond with the current membership list
	return &pb.JoinResponse{Membership: s.membership}, nil
}

func main() {
	// Read configuration from environment variables.
	nodeID := os.Getenv("NODE_ID")
	if nodeID == "" {
		log.Fatal("NODE_ID not set")
	}
	membersEnv := os.Getenv("MEMBERS")
	if membersEnv == "" {
		log.Fatal("MEMBERS not set")
	}
	// Expect MEMBERS as a comma-separated list, e.g., "node1:50061,node2:50062,node3:50063"
	membership := strings.Split(membersEnv, ",")

	// Compute the port for the Dissemination component.
	// For example, use port 50060 + nodeID (if nodeID is numeric).
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
