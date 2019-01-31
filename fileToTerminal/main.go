package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Program to read the contents of a file from disk and print its contents to the std out

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("Please provide a path to a file")
		os.Exit(-1)
	}
	fp := args[1]
	bs, err := ioutil.ReadFile(fp)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(string(bs))

}
