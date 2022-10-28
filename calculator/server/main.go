package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/marcoslhc/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
)

var (
	addr = flag.String("address", "0.0.0.0", "Server Address")
	port = flag.Int64("port", 5002, "Server Port")
)

type Server struct {
	pb.CalculatorServer
}

func main() {
	flag.Parse()
	serverAddress := fmt.Sprintf("%s:%d", *addr, *port)

	log.Printf("Starting Server in %s", serverAddress)

	lis, err := net.Listen("tcp", serverAddress)

	if err != nil {
		log.Fatalf("Can't listen to %s: %v", serverAddress, err)
	}

	log.Printf("Listening in %s", serverAddress)

	defer lis.Close()

	s := grpc.NewServer()

	pb.RegisterCalculatorServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Can't serve at %s: %v", serverAddress, err)
	}

}
