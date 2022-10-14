package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

type Handler struct {
	story Story
}

var defaultHandlerTmpl = `
`

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}

func NewHandler(story Story) http.Handler {
	return Handler{story}
}

func (handler Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	path := strings.TrimSpace(req.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]

	if chapter, ok := handler.story[path]; ok {
		err := tpl.Execute(res, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(res, "Something went wrong..", http.StatusInternalServerError)
		}
		return
	}

	http.Error(res, "Chapter not found.", http.StatusNotFound)
}

func JsonStory(reader io.Reader) (Story, error) {
	decoder := json.NewDecoder(reader)
	var story Story
	err := decoder.Decode(&story)
	if err != nil {
		return nil, err
	}
	return story, nil
}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
