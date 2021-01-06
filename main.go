package main

import (
	"html/template"
	"io"
	"net/http"

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
}

// Render ...
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// Hello ...
func Hello(c echo.Context) error {
	m := &Message{
		Title: "Hello world",
		Text:  "hogehoge",
	}
	return c.Render(http.StatusOK, "hello", m)
}

func main() {
	e := echo.New()
	e.GET("/hello", Hello)

	t := &Template{
		templates: template.Must(template.ParseGlob("./*.tpl")),
	}
	e.Renderer = t

	e.Logger.Fatal(e.Start(":8080"))
}
