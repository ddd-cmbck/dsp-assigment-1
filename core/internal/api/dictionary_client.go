package api

import (
	"context"
	"log"
	"time"

	pb "github.com/ddd-cmbck/dsp-assigment-1/proto/dictionary"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DictionaryClient struct {
	client pb.DictionaryClient
}

func NewDictClient(addr string) (*DictionaryClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("[Core Service]Failed to connect to gRPC dictionary server")
	}

	c := pb.NewDictionaryClient(conn)

	return &DictionaryClient{client: c}, nil
}

func (dc *DictionaryClient) WordExists(ctx context.Context, word string) (bool, error) {
	req := &pb.WordRequest{Word: word}

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	resp, err := dc.client.CheckWord(ctx, req)
	if err != nil {
		log.Fatalf("[Dictionary Client] gRPC error: %v", err)
		return false, err
	}

	return resp.Exists, nil
}
