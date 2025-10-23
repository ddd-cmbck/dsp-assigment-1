package api

import (
	"context"
	"fmt"
	"time"

	pb "github.com/ddd-cmbck/dsp-assigment-1/proto/letters"

	"google.golang.org/grpc"
)

type CoreRepo struct {
	client pb.CoreClient
}

// NewCoreRepo connects to the Core gRPC service
func NewCoreRepo(address string) (*CoreRepo, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to core service: %w", err)
	}

	client := pb.NewCoreClient(conn)
	return &CoreRepo{client: client}, nil
}

// GetLetters fetches letters and the center letter from the Core service
func (r *CoreRepo) GetLetters() ([]string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	resp, err := r.client.GetLetters(ctx, &pb.LettersRequest{})
	if err != nil {
		return nil, "", fmt.Errorf("error fetching letters: %w", err)
	}

	return resp.Letters, resp.Center, nil
}
