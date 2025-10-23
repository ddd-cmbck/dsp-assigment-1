package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ddd-cmbck/dsp-assigment-1/client/internal/api"
	"github.com/ddd-cmbck/dsp-assigment-1/client/internal/util"
)

var (
	server_link = flag.String("serverLink", "localhost:4000", "The gRPC server port link")
)

func main() {
	flag.Parse()
	RUN := true
	score := 0
	points := 0

	client, err := api.NewClient(*server_link)
	if err != nil {
		log.Fatalf("Failed to connect to core service: %v", err)
	}

	letters, center, err := client.GetLetters()
	if err != nil {
		log.Fatalf("Failed to get letters: %v", err)
	}

	letters, err = util.RenderWord(letters, center)
	if err != nil {
		log.Fatalf("Failed to render letters: %v", err)
	}

	fmt.Print(
		"Spelling Bee!\n",
		"Enter \\qt or press CTRL + C if you want to close the game\n\n",
	)

	for RUN {

		fmt.Print(letters, "\n")

		var word string
		fmt.Print("> ")
		fmt.Scanln(&word)

		if word == "\\qt" {
			break
		}

		fmt.Println("Send the word back to core service...")
		// points, err = repo.sendWord(word)

		points = 10
		score += points
		fmt.Printf("Valid word scoring %d points. Current score: %d  \n\n", points, score)

	}

	fmt.Println("Thank you for playing!! >:)")
}
