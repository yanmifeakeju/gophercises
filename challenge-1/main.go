package main

import (
	"log"
	"os"
	"quiz/quiz"
)

func main() {
	source := "problems.csv"
	f, err := os.Open(source)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_ = quiz.New(f)

}
