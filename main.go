package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Run Test Server ...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
