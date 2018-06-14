package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/saboyutaka/grpc-gw-sample/gwsamplepb"
	"strings"
)

const (
	port = ":9090"
)

type server struct{}

func (s *server) Echo(ctx context.Context, in *pb.RequestMessage) (*pb.ResponseMessage, error) {
	return &pb.ResponseMessage{Result: strings.ToUpper(in.GetValue())}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGatewaySampleServer(s, &server{})
	// Register reflection service on gRPC server.
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
