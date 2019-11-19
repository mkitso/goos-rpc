package main

import (
	"fmt"
	"log"
	"time"

	pb "github.com/mkitso/goos-rpc/os"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	os := pb.NewOSClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println("Starting...")

	r, err := os.Hostname(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("%T %v", err, err)
	}

	_, err = os.Mkdir(ctx, &pb.MkdirInput{P: "/tmp/fooBar.lala", M: 0777})
	if err != nil {
		log.Fatalf("%T %v", err, err)
	}

	stat, err := os.Stat(ctx, &pb.SingleString{S: "/tmp/fooBar.lala"})
	if err != nil {
		log.Fatalf("%T %v", err, err)
	}
	fmt.Println(stat)

	_, err = os.MkdirAll(ctx, &pb.MkdirInput{P: "/tmp/a/b/c/d/e/fooBar.lala", M: 0777})
	if err != nil {
		log.Fatalf("%T %v", err, err)
	}

	_, err = os.Remove(ctx, &pb.SingleString{S: "/tmp/fooBar.lala"})
	if err != nil {
		log.Fatalf("%T %v", err, err)
	}

	_, err = os.RemoveAll(ctx, &pb.SingleString{S: "/tmp/a"})
	if err != nil {
		log.Fatalf("%T %v", err, err)
	}

	fmt.Printf("%s", r.S)
}
