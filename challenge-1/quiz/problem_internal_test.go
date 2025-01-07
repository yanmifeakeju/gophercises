package quiz

import "testing"

func TestToProblemsFromValidData(t *testing.T) {
	data := [][]string{
		{"What is 2+2?", "4"},
		{"Capital of France?", "Paris"},
	}

	problems := toProblems(data)

	if len(problems) != 2 {
		t.Errorf("Expected 2 problems, got %d", len(problems))
	}

	if problems[0].question != "What is 2+2?" || problems[0].answer != "4" {
		t.Errorf("First problem does not match expected values")
	}

	if problems[1].question != "Capital of France?" || problems[1].answer != "Paris" {
		t.Errorf("Second problem does not match expected values")
	}
}

func TestToProblemsWithNilInput(t *testing.T) {
	var data [][]string = nil

	problems := toProblems(data)

	if len(problems) != 0 {
		t.Errorf("Expected empty problems array for nil input, got length %d", len(problems))
	}
}
