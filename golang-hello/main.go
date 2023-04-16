
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World! I'm Golang by Hilmi")
	})

	http.ListenAndServe(":8080", nil)
}
