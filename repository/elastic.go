package repository

import (
	// "encoding/json"
	"fmt"
	"github.com/belogik/goes"
	elastigo "github.com/mattbaird/elastigo/lib"
	"strconv"
	"strings"
)

var (
	ES_HOST = "localhost"
	ES_PORT = "9200"
)

func Search() {
	conn := goes.NewConnection(ES_HOST, ES_PORT)

	docType := "note"
	indexName := "sukimono"

	queyr := map[string]interface{}{
		"query": map[string]interface{}{
			"filtered": map[string]interface{}{
				"filter": map[string]interface{}{
				// "term": map[string]interface{}{
				// 	"uuid": "9b46b0dd3a8083c070c3b9953bb5f3f95c5ab4da",
				// },
				},
			},
		},
	}

	res, _ := conn.Search(queyr, []string{indexName}, []string{docType}, nil)

	for _, v := range res.Hits.Hits {
		fmt.Println(v.Source)
		mup := v.Source
		var note Note
		note.Id = int(mup["id"].(float64))
		note.UUId = mup["uuid"].(string)
		note.Name = mup["name"].(string)
		note.Content = mup["content"].(string)
		note.Tags = strings.Split(mup["tags"].(string), ",")
		note.Created = int64(mup["created"].(float64))
		fmt.Println(note)
	}
}

func Index(note *Note) {
	fmt.Println("elastic")

	indexName := "sukimono"

	conn := goes.NewConnection(ES_HOST, ES_PORT)
	// defer conn.DeleteIndex(indexName)

	fmt.Println(note)

	docType := "note"
	docId := strconv.Itoa(note.Id)

	d := goes.Document{
		Index: indexName,
		Type:  docType,
		Id:    docId,
		Fields: map[string]interface{}{
			"id":      note.Id,
			"uuid":    note.UUId,
			"name":    note.Name,
			"content": note.Content,
			"tags":    strings.Join(note.Tags, ","),
			"created": note.Created,
		},
	}

	conn.Index(d, nil)
}

func DeleteIndex() {
	indexName := "sukimono"
	conn := goes.NewConnection(ES_HOST, ES_PORT)
	defer conn.DeleteIndex(indexName)
}
