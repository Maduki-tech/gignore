package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	//https://raw.githubusercontent.com/github/gitignore/main/AL.gitignore
	p := tea.NewProgram(initialModel())
	_, err := p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting program: %v", err)
		os.Exit(1)
	}

}
