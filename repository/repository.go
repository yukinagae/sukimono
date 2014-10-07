package repository

import (
	"fmt"
)

type Note struct {
	Id      int
	Name    string
	Content string
	Tags    []int
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
		Name:    name,
		Content: content,
		Tags:    make([]int, 0),
	}
}

func (repo *Repository) Insert(n Note) *Note {
	n.Id = repo.Id
	repo.notes = append(repo.notes, &n)
	repo.Id = repo.Id + 1
	return &n
}

func (repo *Repository) Update(n Note) *Note {
	note := repo.Select(n.Id)
	note.Name = n.Name
	note.Content = n.Content
	note.Tags = n.Tags
	return note
}

func (repo *Repository) Save(n Note) *Note {
	note := repo.Select(n.Id)
	if note == nil {
		return repo.Insert(n)
	} else {
		return repo.Update(n)
	}
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

// func main() {
// 	repo := NewRepo()
// 	n1 := NewNote("name1", "content1")
// 	repo.Insert(n1)
// 	n2 := NewNote("name2", "content1")
// 	n2.Id = 0
// 	repo.Update(n2)
// 	repo.Dump()
// }
