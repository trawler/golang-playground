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

	fmt.Printf("serving pages on port %s...\n", listenPort)
	http.HandleFunc("/", handler)

	// Parse config file
	site, err := parseYamlFile("./config.yaml")
	if err != nil {
		log.Printf("Error parsing config file: %s", err)
	}

	for _, site := range site.Site {
		if site.User != "" {
			// Use SimpleBasicAuth from: https://girishjoshi.io/post/implementing-http-basic-authentication-in-golang/
			authHandler := httpauth.SimpleBasicAuth(site.User, "password")(http.HandlerFunc(handler))
			http.Handle(site.Path, authHandler)
		}
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", listenPort), nil))
}
