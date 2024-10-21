package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "google.golang.org/grpc/examples/features/proto/echo"
)

var useProxy = flag.Bool("proxy", false, "")

func main() {
	flag.Parse()

	addr := "localhost:50052"
	if *useProxy {
		addr = "localhost:50053"
	}

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	go func() {
		time.Sleep(3 * time.Second)
		conn.Close()
	}()

	c := pb.NewEchoClient(conn)
	fmt.Println("Performing unary request")
	res, err := c.UnaryEcho(context.Background(), &pb.EchoRequest{Message: "keepalive demo"})
	if err != nil {
		log.Fatalf("unexpected error from UnaryEcho: %v", err)
	}
	fmt.Println("RPC response:", res)
}
