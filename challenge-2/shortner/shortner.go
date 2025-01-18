package shortner

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p, ok := pathsToUrls[r.URL.Path]
		if !ok {
			fallback.ServeHTTP(w, r)
			return
		}

		http.Redirect(w, r, p, http.StatusSeeOther)
	}

}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

type urlPath struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var paths []urlPath
	err := yaml.Unmarshal(yml, &paths)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling YAML: %w", err)
	}

	pathUrls := make(map[string]string)

	for _, v := range paths {
		pathUrls[v.Path] = v.Url
	}

	return MapHandler(pathUrls, fallback), nil

}
