package main

import (
	"net/http"

	"github.com/nadama95/godaisyui/cmd/page"
)

func main() {

	http.HandleFunc("/", render)

	http.ListenAndServe(":3000", nil)
}

func render(w http.ResponseWriter, r *http.Request) {
	cmp := page.TestPage()
	cmp.Render(r.Context(), w)
}
