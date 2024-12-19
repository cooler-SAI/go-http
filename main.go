package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Hello go-http")
	fmt.Println(time.Now())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fprintf, err := fmt.Fprintf(w, "Hello go-http")
		if err != nil {
			fmt.Println(fprintf)
			return
		}
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fprintln2, err := fmt.Fprintln(w, "This is the About page of the Simple Go App.")
		if err != nil {
			fmt.Println(fprintln2)
			return
		}
	})

}
