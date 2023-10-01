package main

// A simple program demonstrating the text area component from the Bubbles
// component library.

import (
	"log/slog"
	"sync"

	communication "github.com/TheBromo/gochat/client/communication"
	terminalview "github.com/TheBromo/gochat/client/terminalview"
	pb "github.com/TheBromo/gochat/common/chat"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	input := make(chan pb.Message)
	output := make(chan pb.Message)
	addr := "localhost:50051"
	userName := "Testname"

	p := tea.NewProgram(terminalview.InitialModel(input, output, userName))

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		communication.ConnectToServer(input, output, addr)
		wg.Done()
	}()

	if _, err := p.Run(); err != nil {
		slog.Error("ui error occured: %s", err)
	}

	wg.Wait()
}
