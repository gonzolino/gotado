package gotado

import (
	"fmt"
	"net/url"
	"path"
)

var baseURL = url.URL{
	Scheme: "https",
	Host:   "my.tado.com",
	Path:   "/api/v2",
}

func apiURL(format string, a ...interface{}) string {
	url := baseURL
	url.Path = path.Join(baseURL.Path, fmt.Sprintf(format, a...))
	return url.String()
}

// func api(path ...string) *url.URL {
// 	u := appendPath(baseURL, path...)
// 	return &u
// }

// func appendPath(u url.URL, pathElems ...string) url.URL {
// 	if u.Path == "" {
// 		u.Path = "/"
// 	}
// 	u.Path = path.Join(append([]string{u.Path}, pathElems...)...)
// 	return u
// }
