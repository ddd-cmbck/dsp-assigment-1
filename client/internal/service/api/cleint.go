package api

import (
	"context"
	"fmt"
	"time"

	pb "github.com/ddd-cmbck/dsp-assigment-1/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CoreRepo struct {
	client pb.CoreClient
}

func NewCoreRepo(address string) (*CoreRepo, error) {
	// Create client connection using NewClient
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	client := pb.NewCoreClient(conn)
	return &CoreRepo{client: client}, nil
}

// GetLetters remains the same
func (r *CoreRepo) GetLetters() ([]string, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	resp, err := r.client.GetLetters(ctx, &pb.LettersRequest{})
	if err != nil {
		return nil, "", fmt.Errorf("error fetching letters: %w", err)
	}

	return resp.Letters, resp.Center, nil
}
