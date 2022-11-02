package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/marcoslhc/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("address", "0.0.0.0", "Server Address")
	port = flag.Int64("port", 5001, "Server Port")
)

type executor struct {
	client pb.GreetServiceClient
}
func (e executor) doGreet(name string) {
	res, err := e.client.Greet(context.Background(), &pb.GreetRequest{
		Name: name,
	})

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	log.Printf("%v", res)
}

func (e executor) doGreetManyTime(name string) {
	res, err := e.client.GreetManyTimes(context.Background(), &pb.GreetRequest{
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

func (e executor) doLongGreet() {
	log.Println("doLongGreet Invoked")
	reqs := []*pb.GreetRequest {
		{Name: "Marcos"},
		{Name: "Jeff"},
		{Name: "Trip"},
		{Name: "Calla"},
	}

	stream, err := e.client.LongGreet(context.Background())

	if err != nil {
		log.Fatal("Error")
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Printf("%v", res)
}
func main() {
	flag.Parse()

	cmd := flag.Arg(0)

	serverAddr := fmt.Sprintf("%s:%d", *addr, *port)
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Fail to connect %v", err)
	}

	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)
	exe := executor{client}
	switch cmd {
	case "once":
		name := flag.Arg(1)

		if name == "" {
			log.Fatal("Name not specified, please pass a --name flag")
		}
		exe.doGreet(name)
	case "many":
		name := flag.Arg(1)

		if name == "" {
			log.Fatal("Name not specified, please pass a --name flag")
		}
		exe.doGreetManyTime(name)
	case "long":
		exe.doLongGreet()
	default:
		log.Fatal("That command don't exist")
	}

}
