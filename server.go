package main

import (
	"encoding/json"
	_ "fmt"
	"github.com/gorilla/mux"
	"github.com/yukinagae/sukimono/repository"
	"net/http"
	"strconv"
)

type Context struct {
	Repo *repository.Repository
}

func (c *Context) Close() {
	if c != nil {
		c.Repo.Close()
	}
}

func NewContext() (*Context, error) {
	return &Context{
		Repo: repository.NewRepo(),
	}, nil
}

var ctx, _ = NewContext()

func Ctx() *Context {
	if ctx == nil {
		c, _ := NewContext()
		return c
	} else {
		return ctx
	}
}

func ListHandler(c *Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		notes := c.Repo.List()
		js, _ := json.Marshal(notes)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func FindHandler(c *Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		param := vars["id"]
		id, _ := strconv.Atoi(param)
		note := c.Repo.Select(id)
		js, _ := json.Marshal(note)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func CreateHandler(c *Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var note repository.Note

		decoder.Decode(&note)
		n := c.Repo.Insert(note)
		js, _ := json.Marshal(n)

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func main() {

	r := mux.NewRouter()

	c := Ctx()
	defer c.Close()

	// TODO dummy data
	n1 := repository.NewNote("name1", "content1")
	c.Repo.Insert(n1)
	n2 := repository.NewNote("name2", "content2")
	c.Repo.Insert(n2)

	// routing
	r.HandleFunc("/api/list", ListHandler(c)).Methods("GET")
	r.HandleFunc("/api/{id}", FindHandler(c)).Methods("GET")
	r.HandleFunc("/api/new", CreateHandler(c)).Methods("POST")

	// static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	// start HTTP server
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}
