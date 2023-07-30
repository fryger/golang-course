package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v3"
)

type Redir struct {
	Path string `yaml:"path" json:"path"`
	Url  string `yaml:"url" json:"url"`
}

func MapHandler(urlsmap map[string]string, mux *http.ServeMux) *http.ServeMux {
	for path, url := range urlsmap {
		mux.Handle(path, http.RedirectHandler(url, http.StatusSeeOther))
	}

	return mux
}

func MapUrls(data []Redir, mux *http.ServeMux) *http.ServeMux {
	urlData := make(map[string]string)

	for _, entry := range data {
		urlData[entry.Path] = entry.Url
	}

	mux = MapHandler(urlData, mux)

	return mux
}

func YAMLHandler(yamldata []byte, mux *http.ServeMux) (*http.ServeMux, error) {

	var data []Redir

	err := yaml.Unmarshal(yamldata, &data)

	mux = MapUrls(data, mux)

	return mux, err
}

func JSONHandler(jsondata []byte, mux *http.ServeMux) (*http.ServeMux, error) {

	var data []Redir

	err := json.Unmarshal(jsondata, &data)

	mux = MapUrls(data, mux)

	return mux, err
}
