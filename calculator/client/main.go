package main

import (
	"flag"
	"log"
	"strconv"

	pb "github.com/marcoslhc/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type executor struct {
	client pb.CalculatorClient
}

func main() {
	flag.Parse()
	cmd := flag.Arg(0)
	conn, err := grpc.Dial("0.0.0.0:5002", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Server not available: %v", err)
	}

	client := pb.NewCalculatorClient(conn)
	exe := &executor{
		client: client,
	}

	switch cmd {
	case "sum":
		a, err := strconv.ParseInt(flag.Arg(1), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		b, err := strconv.ParseInt(flag.Arg(2), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		exe.doSum(a, b)
	case "primes":
		prime, err:=strconv.ParseInt(flag.Arg(1), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		exe.doPrimes(prime)
	case "average":
		exe.doAverage()
	default:
		log.Fatal("That command don't exists")
	}
}
