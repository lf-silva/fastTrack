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
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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
	fmt.Printf("You got %d correct answers!\n", r.CorrectAnswers)
	fmt.Printf("You were better than %d%% of all quizzers", r.Score)
}
