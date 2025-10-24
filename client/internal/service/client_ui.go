package service

import "fmt"

func PrintMessage(points int32) {
	if points == 0 {
		fmt.Print("Sorry, that is not a valid word. ")
	}

	fmt.Printf("Valid word scoring %d points. ", points)

}

func PrintScore(score int32) {
	fmt.Printf("Current score is %d\n\n", score)
}
