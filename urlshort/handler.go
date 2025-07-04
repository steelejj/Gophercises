package urlshort

import (
	"fmt"
	"log/slog"
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
	//	TODO: Implement this...
	slog.Info("running map handling: ", slog.Any("paths", pathsToUrls))
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info(r.URL.Path)
		targetPath, ok := pathsToUrls[r.URL.Path]
		if ok {
			http.Redirect(w, r, targetPath, 301) // todo - this should really be a constant
		} else {
			fallback.ServeHTTP(w, r)
		}
		return
	}
}

type YamlPathURL struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
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
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var yamlData []YamlPathURL
	err := yaml.Unmarshal(yml, &yamlData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse yaml data")
	}

	slog.Info("running map handling: ", slog.Any("paths", yamlData))
	return func(w http.ResponseWriter, r *http.Request) {
		for _, redirectMap := range yamlData {
			if redirectMap.Path == r.URL.Path {
				http.Redirect(w, r, redirectMap.Url, 301)
				break
			}
		}
		fallback.ServeHTTP(w, r)
		return

	}, nil
}
