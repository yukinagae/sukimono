package repository

import (
	"encoding/json"
	"fmt"
	elastigo "github.com/mattbaird/elastigo/lib"
	"strconv"
)

var ID = 0

func Search(query string) []Note {
	conn := Connect()

	out, _ := conn.Search("sukimono", "note", nil, "")

	var notes []Note

	for _, v := range out.Hits.Hits {
		var raw *json.RawMessage
		raw = v.Source
		j, err := json.Marshal(&raw)
		if err != nil {
			panic(err)
		}
		var n Note
		json.Unmarshal(j, &n)
		id, _ := strconv.Atoi(v.Id)
		n.Id = id
		notes = append(notes, n)
	}
	fmt.Println(notes)
	return notes
}

func Connect() *elastigo.Conn {
	conn := elastigo.NewConn()
	fmt.Println(conn)
	conn.Domain = "localhost"
	conn.Port = "9200"
	return conn
}

func Index(note Note) Note {
	conn := Connect()
	v, _ := conn.Index("sukimono", "note", strconv.Itoa(ID), nil, &note)
	ID = ID + 1
	fmt.Println(v)
	return note
}

func DeleteIndex() {
	conn := Connect()
	defer conn.DeleteIndex("sukimono")
}
