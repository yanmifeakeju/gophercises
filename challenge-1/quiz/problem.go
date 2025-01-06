package quiz

type problem struct {
	question string
	answer   string
}

func createProblems(data [][]string) []problem {
	problems := make([]problem, len(data))
	for i, d := range data {
		problems[i] = problem{question: d[0], answer: d[1]}
	}

	return problems
}
