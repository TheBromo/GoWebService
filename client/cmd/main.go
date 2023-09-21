package main

// A simple program demonstrating the text area component from the Bubbles
// component library.

import (
	"log"

	terminalview "github.com/TheBromo/gochat/client/terminalview"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(terminalview.InitialModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
