package repository

import (
	_ "fmt"
)

type Note struct {
	Id      int
	Name    string
	Content string
	Tags    []int
}

type Repository struct {
	Id    int
	notes []Note
}

func NewRepo() *Repository {
	return &Repository{
		Id:    0,
		notes: make([]Note, 0),
	}
}

func NewNote(name string, content string) Note {
	return Note{
		Id:      0,
		Name:    name,
		Content: content,
		Tags:    make([]int, 0),
	}
}

func (repo *Repository) Insert(n Note) Note {
	n.Id = repo.Id
	repo.notes = append(repo.notes, n)
	repo.Id = repo.Id + 1
	return n
}

func (repo *Repository) List() []Note {
	return repo.notes
}

func (repo *Repository) Select(id int) Note {
	var note Note
	for _, value := range repo.List() {
		if value.Id == id {
			return value
		}
	}
	return note
}

func (repo *Repository) Close() {
	// TODO
}
