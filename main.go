package main

import (
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// Template ...
type Template struct {
	templates *template.Template
}

// Message ...
type Message struct {
	Title string
	Text  string
	ID    int
}

// Render ...
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// articleIndex ...
func articleIndex(c echo.Context) error {
	data := &Message{
		Title: "Article Index",
		Text:  "index!",
	}
	return c.Render(http.StatusOK, "article/index.html.tpl", data)
}

// articleNew
func articleNew(c echo.Context) error {
	data := &Message{
		Title: "Article New",
		Text:  "New!",
	}

	return c.Render(http.StatusOK, "article/new.html.tpl", data)
}

// articleShow
func articleShow(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := &Message{
		Title: "Article Show",
		Text:  "Show!",
		ID:    id,
	}

	return c.Render(http.StatusOK, "article/show.html.tpl", data)
}

// articleEdit ...
func articleEdit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := &Message{
		Title: "Article Edit",
		Text:  "Edit!",
		ID:    id,
	}

	return c.Render(http.StatusOK, "article/edit.html.tpl", data)
}

func main() {
	e := echo.New()
	e.GET("/hello", articleIndex)
	e.GET("/new", articleNew)
	e.GET("/:id", articleShow)
	e.GET("/edit/:id", articleEdit)

	t := &Template{
		templates: template.Must(template.ParseGlob("./*.tpl")),
	}
	e.Renderer = t

	e.Logger.Fatal(e.Start(":8080"))
}
