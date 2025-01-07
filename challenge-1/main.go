package main

import (
	"flag"
	"log"
	"os"
	"quiz/quiz"
)

func main() {
	source := flag.String("source", "problems.csv", "csv file in format question,answer.")
	flag.Parse()

	f, err := os.Open(*source)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	q := quiz.New(f)

	q.Start(os.Stdin)
}
