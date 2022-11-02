package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/marcoslhc/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
)

var (
	_ pb.GreetServiceServer = &Server{}
)

type Server struct {
	pb.GreetServiceServer
}

var (
	addr = flag.String("address", "0.0.0.0", "Server Address")
	port = flag.Int64("port", 5001, "Server Port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *addr, *port))

	if err != nil {
		log.Fatalf("failed to listen on: %v\n", err)
	}

	defer lis.Close()

	log.Printf("listening on %s", fmt.Sprintf("%s:%d", *addr, *port))

	// opts := []grpc.ServerOption{}

	// opts = append(opts, grpc.Creds(local.NewCredentials()))

	s := grpc.NewServer()

	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serv %v\n", err)
	}
}
