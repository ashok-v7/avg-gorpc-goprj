package main

import (
	"log"

	pb "github.com/tknoptn/avg-stuff-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}

	defer conn.Close()
	client := pb.NewMeetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"avg", "ALice", "Bob", "CHarlie"},
	}

	//unary api method
	//	callSayHello(client)

	//server stream
	//	callHelloServerStream(client, names)

	//clientstream
	//callSayHelloClientStream(client, names)

	callSayHelloBidirectonalStream(client, names)

}
