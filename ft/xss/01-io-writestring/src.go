// Issue 89
// Writting tainted data into http.ResponseWriter can
// result in Cross-Site Scripting.

package main

import (
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Query().Get("param"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":5000", nil)
}
