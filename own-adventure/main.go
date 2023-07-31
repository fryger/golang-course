//https://courses.calhoun.io/lessons/les_goph_06

package main

import (
	"encoding/json"
	"fmt"
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

var router *chi.Mux

func main() {
	stories := readStory(filename)

	fmt.Print(stories)
}
