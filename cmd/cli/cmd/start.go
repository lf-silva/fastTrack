/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lf-silva/fastTrack/internal/model"
	"github.com/lf-silva/fastTrack/internal/ui/multiSelect"
	"github.com/lf-silva/fastTrack/internal/ui/program"
	"github.com/spf13/cobra"
)

var (
	logo = `
▗▄▄▄▖ ▗▄▖  ▗▄▄▖▗▄▄▄▖    ▗▄▄▄▖▗▄▄▖  ▗▄▖  ▗▄▄▖▗▖ ▗▖    ▗▄▄▄▖ ▗▖ ▗▖▗▄▄▄▖▗▄▄▄▄▖
▐▌   ▐▌ ▐▌▐▌     █        █  ▐▌ ▐▌▐▌ ▐▌▐▌   ▐▌▗▞▘    ▐▌ ▐▌ ▐▌ ▐▌  █     ▗▞▘
▐▛▀▀▘▐▛▀▜▌ ▝▀▚▖  █        █  ▐▛▀▚▖▐▛▀▜▌▐▌   ▐▛▚▖     ▐▌ ▐▌ ▐▌ ▐▌  █   ▗▞▘  
▐▌   ▐▌ ▐▌▗▄▄▞▘  █        █  ▐▌ ▐▌▐▌ ▐▌▝▚▄▄▖▐▌ ▐▌    ▐▙▄▟▙▖▝▚▄▞▘▗▄█▄▖▐▙▄▄▄▖
`
	logoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#257FAD")).Bold(true)
	popStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start your quiz!",
	Long: `Start quiz will display questions, one by one, with 4 possible answers.
You can choose one answer per question and you should press 'Enter' or 'space'
to retrieve the next question.
After answering the last question, the program will display both the total number
of correct answers and do you compare to others.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", logoStyle.Render(logo))

		questions := getQuestions()

		program := &program.Project{}
		var answers []model.Answer
		for _, q := range questions {
			p := tea.NewProgram(multiSelect.InitialModel(q.Question, q.Answers, program))
			result, err := p.Run()
			if err != nil {
				fmt.Printf("Error: %v", err)
				os.Exit(1)
			}
			program.ExitCLI(p)

			s := result.(multiSelect.Model)

			answers = append(answers, model.Answer{QuestionID: q.ID, UserAnswer: s.GetCursor()})
		}
		result := submitResult(answers)
		showResult(result)
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func getQuestions() []model.Question {
	resp, err := http.Get("http://localhost:5000/questions")
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer resp.Body.Close()

	var questions []model.Question
	if err := json.NewDecoder(resp.Body).Decode(&questions); err != nil {
		fmt.Println("unable to decode questions")
		return nil
	}
	return questions
}

func submitResult(answers []model.Answer) model.Result {
	body, err := json.Marshal(answers)
	if err != nil {
		fmt.Println("error when marshall")
		os.Exit(1)
	}
	resp, err := http.Post("http://localhost:5000/submit", "application/json", bytes.NewReader([]byte(body)))
	if err != nil {
		fmt.Println("error on submit")
		os.Exit(1)
	}
	var result model.Result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("unable to decode questions")
		os.Exit(1)
	}
	return result
}

func showResult(r model.Result) {
	a := popStyle.Render(strconv.Itoa(r.CorrectAnswers))
	s := popStyle.Render(strconv.Itoa(r.Score) + "%")

	fmt.Printf("You got %s correct answers!\n", a)
	fmt.Printf("You were better than %s of all quizzers", s)
}
