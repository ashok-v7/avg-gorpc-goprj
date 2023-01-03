package main

/*
call the server from client
send the request with names and get back the response as a stream
and print it out
*/

import (
	"context"
	"io"
	"log"

	pb "github.com/tknoptn/avg-stuff-grpc/proto"
)

func callHelloServerStream(client pb.MeetServiceClient, names *pb.NamesList) {
	log.Printf("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("could not sent names %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while streaming %v", err)
		}
		log.Println(message)
	}
	log.Printf("STREAMING FINISHED")
}
