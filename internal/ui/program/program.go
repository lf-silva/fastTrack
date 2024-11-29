package program

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Project struct {
	Exit bool
}

func (p *Project) ExitCLI(tprogram *tea.Program) {
	if p.Exit {
		// logo render here
		if err := tprogram.ReleaseTerminal(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		os.Exit(0)
	}
}
