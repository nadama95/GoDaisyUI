package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/nadama95/godaisyui/cmd/page"
)

func main() {

	http.Handle("/", templ.Handler(page.TestPage()))

	http.ListenAndServe(":3000", nil)
}
