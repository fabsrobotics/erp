package main

///////////////
//  Imports  //
///////////////

import (
	// Golang
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func ServeFile(w http.ResponseWriter, r *http.Request, path string) {
	safePath, err := url.QueryUnescape(path)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	http.ServeFile(w, r, safePath)
}

func GetRouter(w http.ResponseWriter, r *http.Request) {
	ServeFile(w, r, "static/"+r.URL.Path)
}

func MainRouter(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		GetRouter(w, r)
	} else {
		w.Write([]byte("Method not permited"))
	}
}

func main() {
	// Initialize Server
	fmt.Println("Server initialized on port " + os.Getenv("PORT"))
	// Main Router
	http.HandleFunc("/", MainRouter)
	// Listen And Serve
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
