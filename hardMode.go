package main

import (
	"fmt"
	"io"
	"os"
)

func readFileData() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}

// to check the result the create one txt file in same dir and run the following cmt go rum main.go hardMode.go <txt file name>
