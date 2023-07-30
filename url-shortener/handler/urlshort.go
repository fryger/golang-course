package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v3"
)

type Redir struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func MapHandler(urlsmap map[string]string, mux *http.ServeMux) *http.ServeMux {
	for path, url := range urlsmap {
		mux.Handle(path, http.RedirectHandler(url, http.StatusSeeOther))
	}

	return mux
}

func YAMLHandler(yamldata []byte, mux *http.ServeMux) (*http.ServeMux, error) {

	var data []Redir

	err := yaml.Unmarshal(yamldata, &data)

	urlData := make(map[string]string)

	for _, entry := range data {
		urlData[entry.Path] = entry.Url
	}

	mux = MapHandler(urlData, mux)

	return mux, err
}
