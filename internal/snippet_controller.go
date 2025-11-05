package internal

import (
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
)

func NewSnippetController(
	snippetService SnippetService,
) SnippetController {
	return SnippetController{
		snippetService,
	}
}

type SnippetController struct {
	snippetService SnippetService
}

func (c SnippetController) Handler(mux *http.ServeMux) {
	mux.HandleFunc("GET /snippets", c.listPage)
	mux.HandleFunc("GET /snippets/{id}", c.showPage)
	mux.HandleFunc("GET /snippets/_/defaultMain", c.htmxDefaultMain)
	mux.HandleFunc("GET /snippets/_/createForm", c.htmxCreateForm)
	mux.HandleFunc("GET /snippets/_/updateForm", c.htmxUpdateForm)
	mux.HandleFunc("GET /snippets/_/updateList", c.htmxUpdateList)

	mux.HandleFunc("POST /snippets/_/create", c.htmxCreate)
	mux.HandleFunc("POST /snippets/_/update", c.htmxUpdate)
	mux.HandleFunc("POST /snippets/_/delete", c.htmxDelete)
}

func (c SnippetController) htmxCreateForm(writer http.ResponseWriter, request *http.Request) {
	templ.Handler(CreateSnippetMain()).ServeHTTP(writer, request)
}

func (c SnippetController) htmxUpdateForm(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")
	snippet, err := c.snippetService.GetSnippet(id)
	if err != nil {
		templ.Handler(ErrorSnippetMain(
			fmt.Sprintf("Unable to get snippet: %s", id),
			err.Error(),
		)).ServeHTTP(writer, request)
		return
	}

	templ.Handler(UpdateSnippetMain(snippet)).ServeHTTP(writer, request)
}

func (c SnippetController) htmxDefaultMain(writer http.ResponseWriter, request *http.Request) {
	templ.Handler(DefaultSnippetMain()).ServeHTTP(writer, request)
}

func (c SnippetController) htmxUpdateList(writer http.ResponseWriter, request *http.Request) {
	snippets, err := c.snippetService.GetSnippets()
	if err != nil {
		templ.Handler(SnippetsPage(snippets, ErrorSnippetMain(
			"Unable to get snippets",
			err.Error(),
		))).ServeHTTP(writer, request)
		return
	}

	templ.Handler(snippetList(snippets)).ServeHTTP(writer, request)
}

func (c SnippetController) htmxCreate(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		templ.Handler(ErrorSnippetMain(
			"Unable to parse snippet form",
			err.Error(),
		)).ServeHTTP(writer, request)
		return
	}

	filename := request.Form.Get("filename")
	contents := request.Form.Get("contents")
	snippet, err := c.snippetService.CreateSnippet(filename, contents)
	if err != nil {
		templ.Handler(ErrorSnippetMain(
			"Unable to create snippet",
			err.Error(),
		)).ServeHTTP(writer, request)
		return
	}

	templ.Handler(UpdateSnippetMain(snippet)).ServeHTTP(writer, request)
}

func (c SnippetController) htmxUpdate(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	snippet_id := request.Form.Get("snippet_id")
	filename := request.Form.Get("filename")
	contents := request.Form.Get("contents")
	log.Printf("Update snippet [%s] with filename [%s] and contents [%s]", snippet_id, filename, contents)
	snippet, err := c.snippetService.UpdateSnippet(snippet_id, filename, contents)
	if err != nil {
		templ.Handler(ErrorSnippetMain(
			fmt.Sprintf("Unable to update snippet: %s", snippet_id),
			err.Error(),
		)).ServeHTTP(writer, request)
		return
	}

	templ.Handler(UpdateSnippetMain(snippet)).ServeHTTP(writer, request)
}

func (c SnippetController) htmxDelete(writer http.ResponseWriter, request *http.Request) {
	id := request.URL.Query().Get("id")

	if err := c.snippetService.DeleteSnippet(id); err != nil {
		templ.Handler(ErrorSnippetMain(
			fmt.Sprintf("Unable to delete snippet: %s", id),
			err.Error(),
		)).ServeHTTP(writer, request)
		return
	}

	snippets, err := c.snippetService.GetSnippets()
	if err != nil {
		templ.Handler(SnippetsPage(snippets, ErrorSnippetMain(
			"Unable to get snippets",
			err.Error(),
		))).ServeHTTP(writer, request)
		return
	}

	templ.Handler(snippetList(snippets)).ServeHTTP(writer, request)
}

func (c SnippetController) listPage(writer http.ResponseWriter, request *http.Request) {
	snippets, err := c.snippetService.GetSnippets()
	if err != nil {
		templ.Handler(SnippetsPage(snippets, ErrorSnippetMain(
			"Unable to get snippets",
			err.Error(),
		))).ServeHTTP(writer, request)
		return
	}

	templ.Handler(SnippetsPage(snippets, nil)).ServeHTTP(writer, request)
}

func (c SnippetController) showPage(writer http.ResponseWriter, request *http.Request) {
	snippets, err := c.snippetService.GetSnippets()
	if err != nil {
		templ.Handler(SnippetsPage(snippets, ErrorSnippetMain(
			"Unable to get snippets",
			err.Error(),
		))).ServeHTTP(writer, request)
		return
	}

	templ.Handler(SnippetsPage(snippets, nil)).ServeHTTP(writer, request)
}
