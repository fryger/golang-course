//https://courses.calhoun.io/lessons/les_goph_06

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

var filename string = "gopher.json"

type Story map[string]Chapter

type Chapter struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options"`
}

type Options struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

func catch(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func readStory(filename string) (stories Story) {

	stories = make(Story)

	data, err := os.ReadFile(filename)
	catch(err)

	err = json.Unmarshal(data, &stories)
	catch(err)
	return
}

func parseStory(story Chapter, w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/index.html")
	catch(err)

	err = t.Execute(w, story)
	catch(err)
}

func main() {
	stories := readStory(filename)
	router := chi.NewRouter()

	for path, story := range stories {
		// Create a new variable to hold the current story for this route
		currentStory := story

		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/intro", http.StatusFound)
		})

		router.Route(fmt.Sprintf("/%s", path), func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				parseStory(currentStory, w, r)
			})
		})
	}

	err := http.ListenAndServe(":8004", router)
	catch(err)
}
