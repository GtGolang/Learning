package handlers

import (
	"fmt"
	"net/http"
	"path/filepath" //this is use for cleaning the path.
)

// serving the index page
func Home(w http.ResponseWriter, r *http.Request) {
	//making sure that the path is /
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//fmt.Fprintf(w, "welcome to root calcularor")
	http.ServeFile(w, r, "./static/index.html")
}

// serving the css file
func ServeCss(w http.ResponseWriter, r *http.Request) {
	//if the path doesn't match the route we show the message not found.
	if r.URL.Path != "/static/style.css" {
		http.NotFound(w, r)
	}
	/*we convert the path from /static/style.css to
	\static\style.css*/
	cleanedPath := filepath.Clean(r.URL.Path)

	/*adding . to the begining of the path to get accse to the current
	machine file system.*/
	filePath := "." + cleanedPath

	//we let golang safely stream the file to the browser
	http.ServeFile(w, r, filePath)
	//write on terminal the file is being served
	fmt.Println("the file" + filePath + " is being served")

}

func ServeImg(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/static/images/images.png" {
		http.NotFound(w, r)
	}
	//we convert the "/Images/" to "\Images\"
	cleanedPath := filepath.Clean(r.URL.Path)

	//we add the "." to "/Images/"
	completePath := "." + cleanedPath

	//we let golang safely stream the file to the browser
	http.ServeFile(w, r, completePath)
	//write on terminal the file is being served
	fmt.Println("the file" + completePath + " is being served")

}
func ServeJs(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/static/js/myjs.js" {
		http.NotFound(w, r)
	}
	//we convert the "/stati/js/myjs.js" to "\static\js\myjs.js"
	convertedPath := filepath.Clean(r.URL.Path)
	//adding the "." to the new path "\static\js\myjs.js"
	newPath := "." + convertedPath
	http.ServeFile(w, r, newPath)
	//write on terminal the file is being served
	fmt.Println("the file" + newPath + " is being served")
}
