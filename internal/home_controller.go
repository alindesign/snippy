package internal

import (
	"net/http"
)

func NewHomeController() HomeController {
	return HomeController{}
}

type HomeController struct {
}

func (c HomeController) Handler(mux *http.ServeMux) {
	mux.HandleFunc("GET /", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, "/snippets", http.StatusFound)
	})
}
