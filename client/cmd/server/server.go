package main

import (
	"fmt"
)

func main() {
	RUN := true
	score := 0
	points := 0

	fmt.Print(
		"Spelling Bee!\n",
		"Enter \\qt or press CTRL + C if you want to close the game\n\n",
	)

	fmt.Println("Requesting letters from core service...")
	// letters, err := repo.getLetters()

	for RUN {

		var letters string = "A B C [D] E F G"
		// if err != nil {
		// log.Fatalf("%s", err)
		// }
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
