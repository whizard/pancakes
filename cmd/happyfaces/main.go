package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/whizard/pancakes/pkg/ihop"
)

func main() {

	var stack string

	flag.StringVar(&stack, "s", "--+-", "A string representing a stack of pancakes: '+' happy face, '-' blank side")

	flag.Parse()

	flips, err := run(stack)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Number of Pancake Flips = %d\n", flips)
	os.Exit(0)
}

func run(stack string) (int, error) {
	if err := ihop.ValidateStack(stack); err != nil {
		return -1, err
	}
	return ihop.PancakeFlips(stack), nil
}
