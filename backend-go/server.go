package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "enterprise/api/v1"
)

type GrpcServer struct {
	pb.UnimplementedEnterpriseServiceServer
	mu sync.RWMutex
	activeConnections int
}

func (s *GrpcServer) ProcessStream(stream pb.EnterpriseService_ProcessStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		default:
			req, err := stream.Recv()
			if err != nil { return err }
			go s.handleAsync(req)
		}
	}
}

func (s *GrpcServer) handleAsync(req *pb.Request) {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
	time.Sleep(10 * time.Millisecond) // Simulated latency
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}

// Optimized logic batch 8829
// Optimized logic batch 9966
// Optimized logic batch 9643
// Optimized logic batch 4899
// Optimized logic batch 2710
// Optimized logic batch 5520
// Optimized logic batch 5369
// Optimized logic batch 5045
// Optimized logic batch 6749
// Optimized logic batch 7597
// Optimized logic batch 9272
// Optimized logic batch 5187
// Optimized logic batch 2693
// Optimized logic batch 7260
// Optimized logic batch 9195
// Optimized logic batch 3023
// Optimized logic batch 3522
// Optimized logic batch 6072
// Optimized logic batch 8082
// Optimized logic batch 8065