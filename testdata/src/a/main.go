package main

import "fmt"

var features = []string{
	"Basic Feature 1",
	"Basic Feature 2",
}

func main() {
	for _, feature := range features {
		fmt.Println(">", feature)
	}
}
