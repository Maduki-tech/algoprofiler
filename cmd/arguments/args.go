package arguments

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type Arguments struct {
	Algorithm string
	Input     string
	Random    bool
	Size      int
}

func NewArguments() *Arguments {
	return &Arguments{}
}

func (a *Arguments) Parse() error {
	flag.StringVar(&a.Algorithm, "a", "", "Algorithm to use (Required)")
	flag.StringVar(&a.Input, "i", "", "Input example to use (Separated by comma)")
	flag.BoolVar(&a.Random, "r", false, "Use random input")
	flag.IntVar(&a.Size, "s", 10, "Size of input (Needs to be used with -r) Default: 10")

	flag.Parse()

	if a.Algorithm == "" {
		return fmt.Errorf("Algorithm is required")
	}
	if a.Input != "" && a.Random {
		return fmt.Errorf("Input and Random cannot be used together")
	}

	return nil
}

func (a *Arguments) GetInput() ([]int, error) {
	if a.Input == "" {
		return []int{}, nil
	}
	returnValue := []int{}
	splitString := strings.Split(a.Input, ",")
	for _, value := range splitString {
		i, err := strconv.Atoi(value)
		if err != nil {
			return returnValue, fmt.Errorf("Invalid input")
		}
		returnValue = append(returnValue, i)
	}
	return returnValue, nil

}

// ParseFlagSet: For testing, allows parsing with a custom FlagSet
func (a *Arguments) parseFlagSet(flagSet *flag.FlagSet, args []string) error {
	// Define flags on the custom FlagSet
	flagSet.StringVar(&a.Algorithm, "a", "", "Algorithm to use (Required)")
	flagSet.StringVar(&a.Input, "i", "", "Input example to use (Separated by comma)")
	flagSet.BoolVar(&a.Random, "r", false, "Use random input")
	flagSet.IntVar(&a.Size, "s", 10, "Size of input (Needs to be used with -r) Default: 10")

	// Parse the provided arguments
	err := flagSet.Parse(args)
	if err != nil {
		return err
	}

	// Custom validation
	if a.Algorithm == "" {
		return fmt.Errorf("Algorithm is required")
	}
	if a.Input != "" && a.Random {
		return fmt.Errorf("Input and Random cannot be used together")
	}

	return nil
}
