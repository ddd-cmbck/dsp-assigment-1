package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/ddd-cmbck/dsp-assigment-1/core/internal/api"
	"github.com/ddd-cmbck/dsp-assigment-1/core/internal/util"
	pb "github.com/ddd-cmbck/dsp-assigment-1/proto/core"

	"google.golang.org/grpc"
)

var (
	PORT            = flag.Int("port", 4000, "The gRPC server port")
	DICTIONARY_ADDR = flag.String("dictionary_addr", "localhost:4050", "Dictionary service address")
)

func main() {
	flag.Parse()

	address := fmt.Sprintf("localhost:%d", *PORT)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("[Core Service] Failed to listen on %s: %v", address, err)
	}

	dictClient, err := api.NewDictClient(*DICTIONARY_ADDR)
	if err != nil {
		log.Fatalf("[Core Service] Failed to connect to dictionary service: %v", err)
	}

	// create a server with injected logic
	server := &api.CoreServer{
		GenerateLetters: util.GenerateWord,
		PickCenter:      util.PickOne,
		Dictionary:      dictClient,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCoreServer(grpcServer, server)

	log.Printf("[Core Service] gRPC server listening on %s", address)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("[Core Service] Failed to serve: %v", err)
	}
}
