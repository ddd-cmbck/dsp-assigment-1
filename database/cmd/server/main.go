package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/ddd-cmbck/dsp-assigment-1/database/internal/api"
	"github.com/ddd-cmbck/dsp-assigment-1/database/internal/service"
	pb "github.com/ddd-cmbck/dsp-assigment-1/proto/dictionary"
	"google.golang.org/grpc"
)

var (
	PORT = flag.Int("port", 4050, "The gRPC DIctionary server port")
)

func main() {
	flag.Parse()

	_, err := service.Load("../words_dictionary.json")
	if err != nil {
		log.Fatalf("[Dict Service] Failed to load dictionary: %v", err)
	}

	address := fmt.Sprintf("localhost:%d", *PORT)
	list, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("[Dict Service] Failed to listen on %s: %v", address, err)
	}

	server := &api.DictionaryServer{}
	grpcServer := grpc.NewServer()
	pb.RegisterDictionaryServer(grpcServer, server)

	log.Printf("[Dict Service] gRPC server listening on %s", address)

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("[Core Service] Failed to serve: %v", err)
	}

}
