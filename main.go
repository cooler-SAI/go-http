package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello go-http")
	for {
		time.Sleep(time.Hour)
	}
}
