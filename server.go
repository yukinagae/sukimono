package main

import (
	"encoding/json"
	_ "fmt"
	"github.com/gorilla/mux"
	"github.com/yukinagae/sukimono/repository"
	"net/http"
	// "strconv"
	"time"
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
		// notes := c.Repo.List()
		notes := repository.Search("")
		js, _ := json.Marshal(notes)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

// func FindHandler(c *Context) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		param := vars["id"]
// 		id, _ := strconv.Atoi(param)
// 		note := c.Repo.Select(id)
// 		js, _ := json.Marshal(note)
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(js)
// 	}
// }

func SaveHandler(c *Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var note repository.Note

		decoder.Decode(&note)
		note.Created = time.Now().Unix()

		// n := c.Repo.Insert(note)
		n := repository.Index(note)
		// repository.Index(n)

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
	// n1 := repository.NewNote("name1", "content1")
	// n1.Tags = []string{"hoge1"}
	// c.Repo.Insert(n1)
	// n2 := repository.NewNote("name2", "content2")
	// n2.Tags = []string{"hoge1", "done"}
	// c.Repo.Insert(n2)

	// routing
	r.HandleFunc("/api/list", ListHandler(c)).Methods("GET")
	// r.HandleFunc("/api/{id}", FindHandler(c)).Methods("GET")
	r.HandleFunc("/api/save", SaveHandler(c)).Methods("POST")

	// static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	// start HTTP server
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}

// func main() {
//
// n1 := repository.NewNote("name1", "content1")
// n1.Tags = []string{"hoge1"}
// n2 := repository.NewNote("name2", "content2")
// n2.Tags = []string{"hoge1", "done"}
// n1a := ctx.Repo.Insert(n1)
// n2a := ctx.Repo.Insert(n2)
// repository.Index(n1a)
// repository.Index(n2a)
// notes := repository.Search("")
// fmt.Println(notes)
// repository.DeleteIndex()
// }
