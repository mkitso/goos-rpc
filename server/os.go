package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/mkitso/goos-rpc/os"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement os.OSServer.
type server struct{}

// Hostname implements os.Hostname
func (s *server) Hostname(ctx context.Context, in *pb.Empty) (*pb.SingleString, error) {
	hostname, err := os.Hostname()
	return &pb.SingleString{S: hostname}, err
}

// Mkdir implements os.Mkdir
func (s *server) Mkdir(ctx context.Context, in *pb.MkdirInput) (*pb.Empty, error) {
	path := in.P
	mode := os.FileMode(in.M)
	err := os.Mkdir(path, mode)
	return &pb.Empty{}, err
}

// MkdirAll implements os.MkdirAll
func (s *server) MkdirAll(ctx context.Context, in *pb.MkdirInput) (*pb.Empty, error) {
	path := in.P
	mode := os.FileMode(in.M)
	err := os.MkdirAll(path, mode)
	return &pb.Empty{}, err
}

// Remove implements os.Remove
func (s *server) Remove(ctx context.Context, in *pb.SingleString) (*pb.Empty, error) {
	path := in.S
	err := os.Remove(path)
	return &pb.Empty{}, err
}

// RemoveAll implements os.RemoveAll
func (s *server) RemoveAll(ctx context.Context, in *pb.SingleString) (*pb.Empty, error) {
	path := in.S
	err := os.RemoveAll(path)
	return &pb.Empty{}, err
}

// Stat implements os.Stat
func (s *server) Stat(ctx context.Context, in *pb.SingleString) (*pb.FileInfo, error) {
	path := in.S
	finfo, err := os.Stat(path)
	info := &pb.FileInfo{Name: finfo.Name(), Size: finfo.Size(), Mode: uint32(finfo.Mode()), IsDir: finfo.IsDir()}
	x := finfo.Sys()
	fmt.Printf("%T %+v\n", x, x)
	return info, err
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOSServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
