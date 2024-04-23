package main

import (
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
)

func main(){
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("There has been an error executing the program: %v", err)
        os.Exit(1)
    }
}
