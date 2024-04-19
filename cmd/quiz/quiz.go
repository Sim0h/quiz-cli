package quiz

import (
	"math/rand"
)

// Question represents a quiz question
type Question struct {
	ID         int
	Question   string
	Answers    []string
	CorrectIdx int
}

func GetQuestions() []Question {
	return []Question{
		{
			Question:   "What is the capital of France?",
			Answers:    []string{"Paris", "London", "Berlin", "Rome"},
			CorrectIdx: 0,
		},
		{
			Question:   "What is the largest planet in our solar system?",
			Answers:    []string{"Jupiter", "Saturn", "Mars", "Earth"},
			CorrectIdx: 0,
		},
		{
			Question:   "Which of these is NOT a primary color??",
			Answers:    []string{"Red", "Yellow", "Orange", "Blue"},
			CorrectIdx: 2,
		},
		{
			Question:   "In which year was the first iPhone released?",
			Answers:    []string{"2000", "2004", "2007", "2010"},
			CorrectIdx: 2,
		},
		{
			Question:   "How many sides does a hexagon have?",
			Answers:    []string{"4", "5", "6", "8"},
			CorrectIdx: 2,
		},
		{
			Question:   "What is the largest ocean on Earth?",
			Answers:    []string{"Atlantic Ocean", "Pacific Ocean", "Indian Ocean", "Arctic Ocean"},
			CorrectIdx: 1,
		},
		{
			Question:   "Which of these is a carnivorous animal?",
			Answers:    []string{"Cow", "Lion", "Zebra", "Horse"},
			CorrectIdx: 1,
		},
		{
			Question:   "What is the name of the structure in Paris known for its iron lattice tower?",
			Answers:    []string{"The Louvre Museum", "The Arc de Triomphe", "Notre Dame", "The Eiffel Tower"},
			CorrectIdx: 3,
		},
	}
}

func IsAnswerCorrect(question Question, answerIdx int) bool {
	return answerIdx == question.CorrectIdx
}

func GetRandomQuestion(questions []Question) []Question {
	for i := len(questions) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		questions[i], questions[j] = questions[j], questions[i]
	}

	return questions
}
