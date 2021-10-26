package main

import (
	"net/http"
	"strings"
)

// https://shogo82148.github.io/blog/2016/04/13/serving-static-files-in-golang/
// 脆弱性があるので注意
func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/") {
			http.ServeFile(w, r, r.URL.Path[1:])
		} else {
			http.NotFound(w, r)
		}
	}))
}
