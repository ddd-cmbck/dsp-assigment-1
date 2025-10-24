package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/ddd-cmbck/dsp-assigment-1/client/internal/api"
	"github.com/ddd-cmbck/dsp-assigment-1/client/internal/service"
	"github.com/ddd-cmbck/dsp-assigment-1/client/internal/util"
)

var (
	CORE_ADDR = flag.String("core_addr", "localhost:4000", "The Core gRPC server port")
)

func main() {
	flag.Parse()
	RUN := true
	var score int32

	client, err := api.NewClient(*CORE_ADDR)
	if err != nil {
		log.Fatalf("Failed to connect to core service: %v", err)
	}

	letters, center, err := client.GetLetters(context.Background())
	if err != nil {
		log.Fatalf("Failed to get letters: %v", err)
	}

	formated_letters, err := util.RenderWord(letters, center)
	if err != nil {
		log.Fatalf("Failed to render letters: %v", err)
	}

	fmt.Print(
		"Spelling Bee!\n",
		"Enter \\qt or press CTRL + C if you want to close the game\n\n",
	)

	for RUN {

		fmt.Print(formated_letters, "\n")

		var word string
		fmt.Print("> ")
		fmt.Scanln(&word)

		if word == "\\qt" {
			break
		}

		points, err := client.GetScore(context.Background(), word, letters)
		if err != nil {
			log.Fatalf("Failed to get score: %v", err)
		}

		score += points
		service.PrintMessage(points)
		service.PrintScore(score)

	}

	fmt.Println("Thank you for playing!! >:)")
}
