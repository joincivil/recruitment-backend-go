package main

import (
	"fmt"
	"net/http"

	myhttp "github.com/joincivil/recruitment-backend-go/pkg/http"
)

func main() {
	http.HandleFunc("/", myhttp.DiffHandler)

	fmt.Println("starting up service on port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("error launching server: err: %v", err)
	}
}
