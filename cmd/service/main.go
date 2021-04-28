package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Server started in port NNNN\n")
	for {
		time.Sleep(5 * time.Second) // it's horrible I know
	}
}
