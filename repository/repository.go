package repository

import (
// "fmt"
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

// func main() {
// 	repo := NewRepo()
// 	// fmt.Println(repo.List())

// 	n1 := NewNote("name1", "content1")
// 	// fmt.Println(n1)
// 	repo.Insert(n1)

// 	n2 := NewNote("name2", "content2")
// 	// fmt.Println(n2)
// 	repo.Insert(n2)

// 	fmt.Println(repo.List())

// 	fmt.Println(repo.Select(1))

// }