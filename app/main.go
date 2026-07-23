package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/health", HealthHandler)
	http.HandleFunc("/version", VersionHandler)
	http.HandleFunc("/time", TimeHandler)

	fmt.Println("Server started on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
