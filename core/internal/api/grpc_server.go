package api

import (
	"context"
	"log"

	"github.com/ddd-cmbck/dsp-assigment-1/core/internal/service"
	pb "github.com/ddd-cmbck/dsp-assigment-1/proto/core"
)

type CoreServer struct {
	pb.UnimplementedCoreServer

	GenerateLetters func() []string
	PickCenter      func([]string) string

	//gRPC clients
	Dictionary *DictionaryClient
}

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

func (s *CoreServer) GetScore(ctx context.Context, req *pb.UserWord) (*pb.Score, error) {
	word := req.Word
	letters := req.Letters
	log.Printf("[Core service] Received word from user %s, with next set of letters %s", word, letters)

	exists, err := s.Dictionary.WordExists(ctx, word)
	if err != nil {
		log.Print("[Core service]Failed to call dictionary service")
		return nil, err
	}

	if !exists {
		log.Printf("[Core service] Word not found: %s", word)
		return &pb.Score{Score: 0}, nil
	}

	score, err := service.EvalScore(word, letters)
	if err != nil {
		log.Printf("[Core service] Failed to calculate score: %v", err)
	}

	log.Printf("[Core Service] Sending response to client: Score=%d", score)

	return &pb.Score{Score: score}, nil

}
