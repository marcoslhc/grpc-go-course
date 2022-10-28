package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	pb "github.com/marcoslhc/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("address", "0.0.0.0", "Server Address")
	port = flag.Int64("port", 5001, "Server Port")
)

func doGreet(c pb.GreetServiceClient, name string) {
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		Name: name,
	})

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	log.Printf("%v", res)
}

func doGreetManyTime(c pb.GreetServiceClient, name string) {
	res, err := c.GreetManyTimes(context.Background(), &pb.GreetRequest{
		Name: name,
	})

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	for {
		msg, err := res.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error reading the stream %v", err)
		}

		log.Printf("%v\n", msg)
	}
}
func main() {
	flag.Parse()

	name := flag.Arg(0)

	if name == "" {
		log.Fatal("Name not specified, please pass a --name flag")
	}

	serverAddr := fmt.Sprintf("%s:%d", *addr, *port)
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Fail to connect %v", err)
	}

	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)

	doGreetManyTime(client, name)

}
