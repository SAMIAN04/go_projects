package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
func main(){
r := mux.NewRouter()
r.HandleFunc("/", serveHome)
log.Fatal(http.ListenAndServe(":8080", r))
}


func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(" <body style='background-color: black;'> <h1 style='color: blue;'>Welcome to the Home Page!</h1> </body>"))
}
