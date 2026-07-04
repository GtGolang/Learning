package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
)

// HomeHandler handles request sent to the root "/" path
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//ensure the path is exactly "/" to avoid matching all sub -paths
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//fmt.Fprint(w,"Welcome to the fundation of something Bg")
	http.ServeFile(w, r, "./static/index.html")
	fmt.Println("running the page index.html")
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello from the Api endpoint!"}`))
}
func CssHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/static/style.css" {
		http.NotFound(w, r)
		fmt.Println("ruta no encontrada")
		return
	}
	fmt.Println("ruta encontrada /static/style.css")

	// 1. Clean the path to prevent security issues (directory traversal)
	// r.URL.Path will be something like "/static/style.css"
	// filepath.Clean() does change "/static/style.css" into "\static\style.css"
	cleanedPath := filepath.Clean(r.URL.Path)

	// 2. Join it with your local folder name.
	// This turns "\static\style.css" into ".\static\style.css"
	filePath := "." + cleanedPath

	// 3. Let Go safely stream the file to the browser
	http.ServeFile(w, r, filePath)
	//http.ServeFile(w,r,"./static/style.css")
	fmt.Println("css served")
}

func ServeImages(w http.ResponseWriter, r *http.Request) {
	cleanedPath := filepath.Clean(r.URL.Path)
	filePath := "." + cleanedPath
	http.ServeFile(w, r, filePath)
	fmt.Println("img served")
}
