package quiz

import "strings"

type problem struct {
	question string
	answer   string
}

func (p problem) isValidAnswer(answer string) bool {
	return strings.EqualFold(strings.TrimSpace(p.answer), strings.TrimSpace(answer))
}

func createProblems(data [][]string) []problem {
	problems := make([]problem, len(data))
	for i, d := range data {
		problems[i] = problem{question: d[0], answer: d[1]}
	}

	return problems
}
