package handlers
import(
	"net/http"
	"fmt"
)
// HomeHandler handles request sent to the root "/" path
func HomeHandler(w http.ResponseWriter, r *http.Request){
	 //ensure the path is exactly "/" to avoid matching all sub -paths
	 if r.URL.Path!="/"{
		http.NotFound(w,r)
		return
	 }
	 //fmt.Fprint(w,"Welcome to the fundation of something Bg")
	 http.ServeFile(w,r,"./static/index.html")
	fmt.Println("running the page index.html")
}


func ApiHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-Type","aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "hello from the Api endpoint!"}`))
}

