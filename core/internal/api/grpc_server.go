package api

import (
	"context"
	"log"

	pb "github.com/ddd-cmbck/dsp-assigment-1/proto"
)

// CoreServer implements the Core gRPC service defined in letters.proto
type CoreServer struct {
	pb.UnimplementedCoreServer

	GenerateLetters func() []string
	PickCenter      func([]string) string
}

// GetLetters implements the gRPC method
func (s *CoreServer) GetLetters(ctx context.Context, req *pb.LettersRequest) (*pb.LettersResponse, error) {
	letters := s.GenerateLetters()
	center := s.PickCenter(letters)

	resp := &pb.LettersResponse{
		Letters: letters,
		Center:  center,
	}

	log.Printf("[Core Service] Sending response to client: Letters=%v | Center=%s", resp.Letters, resp.Center)

	return resp, nil
}
