package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello go-http")
	fmt.Println(time.Now())
	fmt.Println("Please press any key to continue....")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

}
