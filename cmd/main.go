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

func runAlgorithm(name string, input []int) {
	switch name {
	case "bubble":
		run(input, name, algorithms.BubbleSort)

	case "merge":
		run(input, name, algorithms.Merge_sort)

	case "quick":
		track := tracker.NewTracker(name)
		timer := tracker.NewTimer(name, track)
		timer.StartTimer()
		algorithms.Quick_sort(input, 0, len(input)-1)
		timer.EndTimer().Seconds()

	default:
		fmt.Println("Invalid algorithm")
		os.Exit(1)
	}
}

type operation func([]int) []int

func run(input []int, name string, op operation) {
	track := tracker.NewTracker(name)
	timer := tracker.NewTimer(name, track)
	timer.StartTimer()

	fmt.Println(op(input))

	timer.EndTimer().Seconds()
}
