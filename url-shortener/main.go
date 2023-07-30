// https://courses.calhoun.io/lessons/les_goph_04

package main

import (
	urlshort "Tutorial/url-shortener/handler"

	_ "github.com/mattn/go-sqlite3"

	"database/sql"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

var filePtr = flag.String("file", "urls.yaml", "a file with map of paths : urls' (default 'urls.yaml')")
var dbPtr = flag.Bool("db", false, "use sqlite db (default: false)")

func main() {
	const dbName string = "urls.db"

	flag.Parse()

	var handler *http.ServeMux

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback

	file, err := os.ReadFile(*filePtr)

	if err != nil {
		panic(err)
	}

	if *dbPtr {
		*filePtr = ".db"
	}

	switch strings.Split(*filePtr, ".")[1] {
	case "yaml":
		handler, err = urlshort.YAMLHandler([]byte(file), mapHandler)

	case "json":
		handler, err = urlshort.JSONHandler([]byte(file), mapHandler)

	case "db":
		db, err := sql.Open("sqlite3", dbName)

		if err != nil {
			panic(err)
		}

		handler, err = urlshort.SQLiteHandler(db, mapHandler)

		if err != nil {
			panic(err)
		}

	default:
		panic("")
	}

	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
