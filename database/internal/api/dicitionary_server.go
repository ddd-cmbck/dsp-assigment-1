package api

import (
	"context"
	"log"

	"github.com/ddd-cmbck/dsp-assigment-1/database/internal/service"
	pb "github.com/ddd-cmbck/dsp-assigment-1/proto/dictionary"
)

type DictionaryServer struct {
	pb.UnimplementedDictionaryServer
}

func (s *DictionaryServer) CheckWord(ctx context.Context, req *pb.WordRequest) (*pb.WordResponse, error) {

	word := req.Word
	log.Printf("[Dictionary Service] Received word %s", word)

	dict, err := service.GetInstance()
	if err != nil {
		log.Println("[Dict Service] Dictionary not initialized:", err)
		return nil, err
	}

	exists, err := dict.CheckWord(word)
	if err != nil {
		log.Printf("[Dictionary service] Error %v", err)
		return &pb.WordResponse{Exists: false}, nil
	}

	log.Printf("[Dict Service] Sending response to client: Exists=%v", exists)

	return &pb.WordResponse{Exists: exists}, nil

}
