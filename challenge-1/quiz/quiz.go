package quiz

import (
	"encoding/csv"
	"io"
	"log"
)

type Quiz struct {
	problems []problem
	score    int
}

func New(source io.Reader) *Quiz {
	reader := csv.NewReader(source)
	data, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	problems := createProblems(data)
	return &Quiz{problems: problems}
}
