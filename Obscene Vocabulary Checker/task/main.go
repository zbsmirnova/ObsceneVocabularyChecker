package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var input string
	fmt.Scan(&input)

	file, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(file))
}
