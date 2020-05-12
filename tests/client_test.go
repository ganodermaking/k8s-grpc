package tests

import (
	"context"
	"time"

	"google.golang.org/grpc"
	pb "k8s-grpc/helloworld"
	"testing"
)

const (
	address = "localhost:32766"
)

func TestSayHello(t *testing.T) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Errorf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	if err != nil {
		t.Errorf("could not greet: %v", err)
	}

	t.Log("Greeting: ", r.GetMessage())
}
