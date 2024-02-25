package main

import (
	"fmt"
	"os"
)

func main() {
	cli := &CLI{
		Stdin: os.Stdin,
		Args:  os.Args[1:],
	}

	out, err := cli.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(out)
	os.Exit(0)
}
