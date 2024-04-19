package cmd

import (
	"fmt"
	"strconv"

	"github.com/Sim0h/quiz-cli/cmd/quiz"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "quiz-cli",
	Short: "A simple quiz application for LIA-application to Fast Track",
	Run: func(cmd *cobra.Command, args []string) {
		// Get quiz questions
		questions := quiz.GetQuestions()
		questions = quiz.GetRandomQuestion(questions)
		score := 0

		for i, question := range questions {
			// Make the Question more readable
			fmt.Printf(color.YellowString("Question %d: %s\n"), i+1, question.Question)

			fmt.Println("Choices:")
			for j, choice := range question.Answers {
				fmt.Printf("%d. %s\n", j+1, choice)
			}

			var userAnswer string
			fmt.Print("Your answer (enter the number): ")
			fmt.Scanln(&userAnswer)

			// Prase input from user, and see if its an acceptable answer
			userChoice, err := strconv.Atoi(userAnswer)
			if err != nil || userChoice < 1 || userChoice > len(question.Answers) {
				fmt.Println("Invalid choice!")
				continue
			}

			// Check if the answer is correct
			if quiz.IsAnswerCorrect(question, userChoice-1) {
				score++
				fmt.Println("Correct!")
			} else {
				fmt.Println("Incorrect!")
			}
		}

		//More color for readability
		fmt.Printf(color.YellowString("Total score: %d/%d\n"), score, len(questions))
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
