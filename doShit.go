package main

import (
	"bufio"
	"os"
)

func main() {
	filename := "silly.go"

	// open output file
	fo, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	// make a write buffer
	w := bufio.NewWriter(fo)
	w.WriteString("package main\n")
	w.Flush()

	//commit it

}
