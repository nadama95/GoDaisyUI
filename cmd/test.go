package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/nadama95/godaisyui/cmd/page"
)

func main() {

	http.Handle("/", templ.Handler(page.TestPage()))

	log.Fatal(http.ListenAndServe(":3000", nil))
}
