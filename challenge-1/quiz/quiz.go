package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
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

	problems := toProblems(data)
	return &Quiz{problems: problems}
}

func (q *Quiz) Start(input io.Reader) {
	fmt.Println("Quiz started:")

	for i := 0; i < len(q.problems); i++ {
		reader := bufio.NewReader(input)
		problem := q.problems[i]
		fmt.Printf("# Problem %d: %s\n", i+1, problem.question)

		answer, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		if problem.isValidAnswer(answer) {
			q.score++
		}
	}

	fmt.Printf("You scored%d out %d\n", q.score, len(q.problems))
}
