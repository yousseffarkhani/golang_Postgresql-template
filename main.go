package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("2")
	fs := http.FileServer(http.Dir(""))
	http.Handle("/", fs)
	http.ListenAndServe(":8080", nil)
}
