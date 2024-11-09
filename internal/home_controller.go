package internal

import (
	"net/http"

	"github.com/a-h/templ"
)

func NewHomeController() HomeController {
	return HomeController{}
}

type HomeController struct {
}

func (c HomeController) Handler(mux *http.ServeMux) {
	mux.HandleFunc("GET /", c.homePage)
}

func (c HomeController) homePage(writer http.ResponseWriter, request *http.Request) {
	component := HomePage()

	templ.Handler(component).ServeHTTP(writer, request)
}
