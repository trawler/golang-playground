package main

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/goji/httpauth"
)

// https://golang.org/doc/articles/wiki/
func handler(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Path
	title := fmt.Sprintf("%s", path.Base(page))
	if page == "/" {
		title = "home"
	}
	fmt.Fprintf(w, "<h1>%s</h1><div>You're now browsing: %s</div>", title, page)
}

func main() {
	listenPort := "8080"

	fmt.Printf("Listening on port %s...\n", listenPort)
	site, err := parseYamlFile("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	for _, site := range site.Site {
		if site.Credential != "" {
			// Use SimpleBasicAuth from: https://girishjoshi.io/post/implementing-http-basic-authentication-in-golang/
			http.Handle(site.Path, httpauth.SimpleBasicAuth(site.Credential, "password")(http.HandlerFunc(handler)))
		} else {
			http.HandleFunc("/", handler)
		}
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", listenPort), nil))
}
