package main

import (
	// "fmt"
	// "github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/yukinagae/sukimono/repository"
	"html/template"
	"net/http"
)

func main() {

	repo := repository.NewRepo()
	defer repo.Close()

	r := mux.NewRouter()

	// TODO dummy data
	n1 := repository.NewNote("name1", "content1")
	repo.Insert(n1)
	n2 := repository.NewNote("name2", "content2")
	repo.Insert(n2)

	// routing
	// GET /
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("index.html")
		notes := repo.List()
		t.Execute(w, notes)
	}).Methods("GET")

	// For static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}
