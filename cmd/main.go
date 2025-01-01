package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/Maduki-tech/algoprofiler/cmd/algorithms"
	"github.com/Maduki-tech/algoprofiler/cmd/arguments"
	"github.com/Maduki-tech/algoprofiler/cmd/tracker"
)

func main() {
	args := arguments.NewArguments()
	err := args.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	input, err := args.GetInput()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(input) == 0 && args.Random {
		input = generateRandomInput(args.Size)
	}

	runAlgorithm(args.Algorithm, input)
}

func generateRandomInput(i int) []int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	min, max := 1, 10000

	slice := make([]int, i)

	for j := range slice {
		slice[j] = rand.Intn(max-min+1) + min
	}

	return slice
}

func runAlgorithm(algorithm string, input []int) {
	switch algorithm {
	case "bubble":
		runBubblesort(input)

	case "merge":

	case "quick":

	default:
		fmt.Println("Invalid algorithm")
		os.Exit(1)
	}
}

func runBubblesort(input []int) {
	track := tracker.NewTracker("BubbleSort")
	timer := tracker.NewTimer("BubbleSort", track)
	timer.StartTimer()

	fmt.Println(algorithms.BubbleSort(input))

	timer.EndTimer().Seconds()
}
