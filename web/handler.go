package web

import (
	"embed"
	"net/http"
)

//go:embed assets/**
var assets embed.FS

func Handler(mux *http.ServeMux) {
	mux.Handle("GET /assets/", http.FileServer(http.FS(assets)))
}
