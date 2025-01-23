package main

import (
	"fmt"
	"os"
	// "net/http"
)

// func handlerFunc(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "<h1>Push it to the limit</h1>")
// }

func main() {
	// http.HandleFunc("/", handlerFunc)
	fmt.Fprintln(os.Stdout, "Hello, World!")
	fmt.Println("Starting the server on :3000...")
	// http.ListenAndServe(":3000", nil)
}
