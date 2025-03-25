package godaisyui

import (
	_ "embed"
	"net/http"
)

//go:embed static/css/godaisyui.css
var CSS []byte

func ServeCss(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	w.Write(CSS)
}
