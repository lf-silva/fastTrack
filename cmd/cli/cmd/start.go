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
	"github.com/lf-silva/fastTrack/internal/model"
	"github.com/lf-silva/fastTrack/internal/ui/multiSelect"
	"github.com/spf13/cobra"
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
		resp, err := http.Get("http://localhost:5000/questions")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer resp.Body.Close()
		var questions []model.Question
		if err := json.NewDecoder(resp.Body).Decode(&questions); err != nil {
			fmt.Println("unable to decode questions")
		}

		var answers []model.Answer
		for _, q := range questions {
			p := tea.NewProgram(multiSelect.InitialModel(q.Question, q.Answers))
			result, err := p.Run()
			if err != nil {
				fmt.Printf("Alas, there's been an error: %v", err)
				os.Exit(1)
			}
			fmt.Println(result)
			s := result.(multiSelect.Model)
			fmt.Println("Cursor is")
			fmt.Println("-----------")
			fmt.Println(s.GetCursor())
			fmt.Println("-----------")
			answers = append(answers, model.Answer{QuestionID: q.ID, UserAnswer: s.GetCursor()})
			//			program.ExitCLI(tprogram)
		}
		fmt.Println("-----------")
		fmt.Println("User Answers")
		fmt.Println(answers)
		fmt.Println("-----------")

		fmt.Println("marshall")
		body, err := json.Marshal(answers)
		if err != nil {
			fmt.Println("error when marshall")
			os.Exit(1)
		}
		fmt.Println("submitting")
		resp, err = http.Post("http://localhost:5000/submit", "application/json", bytes.NewReader([]byte(body)))
		if err != nil {
			fmt.Println("error on submit")
			os.Exit(1)
		}
		fmt.Println("-----------")
		var result model.Result
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			fmt.Println("unable to decode questions")
		}
		fmt.Println(result)
		fmt.Println("-----------")

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
