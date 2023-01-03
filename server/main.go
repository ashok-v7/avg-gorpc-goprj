package main

import (
	"log"
	"net"

	pb "github.com/tknoptn/avg-stuff-grpc/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.MeetServiceServer
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to strat the server %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMeetServiceServer(grpcServer, &helloServer{})
	log.Printf("server started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start : %v", err)
	}

}
