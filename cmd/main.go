package main

import (
	"fmt"
	"os"

	"github.com/Maduki-tech/algoprofiler/cmd/arguments"
)

func main() {
	args := arguments.NewArguments()
	err := args.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch args.Algorithm {
	case "bubble":

	case "merge":

	case "quick":

	default:
		fmt.Println("Invalid algorithm")
		os.Exit(1)
	}

	fmt.Println(args.Algorithm)

}
