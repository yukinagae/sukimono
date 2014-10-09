package repository

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"time"
)

type Note struct {
	Id      int      `json:"id"`
	UUId    string   `json:"uuid"`
	Name    string   `json:"name"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
	Created int64    `json:"created"`
}

type Repository struct {
	Id    int
	notes []*Note
}

func NewRepo() *Repository {
	notes := make([]*Note, 0)
	return &Repository{
		Id:    1,
		notes: notes,
	}
}

func NewNote(name string, content string) Note {
	return Note{
		Id:      1,
		UUId:    createUUID(name),
		Name:    name,
		Content: content,
		Tags:    []string{},
		Created: time.Now().Unix(),
	}
}

func (repo *Repository) Insert(n Note) *Note {
	n.Id = repo.Id
	repo.notes = append(repo.notes, &n)
	repo.Id = repo.Id + 1
	return &n
}

func (repo *Repository) List() []*Note {
	return repo.notes
}

func (repo *Repository) Select(id int) *Note {
	var note *Note
	for _, value := range repo.List() {
		if value.Id == id {
			return value
		}
	}
	return note
}

func (repo *Repository) Close() {
	repo = nil
}

// TODO test purpose
func (repo *Repository) Dump() {
	for _, value := range repo.List() {
		fmt.Println(value)
	}
}

func createUUID(name string) string {

	hasher := sha1.New()
	hasher.Write([]byte(name))
	sha := hasher.Sum(nil)
	s := hex.EncodeToString(sha)
	return s
}
