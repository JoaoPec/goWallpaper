package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"os/exec"
)

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
	tea.Model
}

func initialModel() model {

	wallpapers, err := os.ReadDir("/home/tiramisu/wallpapers/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//Get the names of the wallpapers
	var choices []string
	for _, wallpaper := range wallpapers {
		choices = append(choices, wallpaper.Name())
	}

	return model{
		choices:  choices,
		selected: make(map[int]struct{}),
	}

}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}

			// Execute feh command with the selected wallpaper
			go func() {
				path := "/home/tiramisu/wallpapers/" + m.choices[m.cursor]
				cmd := exec.Command("feh", "--bg-scale", path)
				err := cmd.Run()
				if err != nil {
					fmt.Println("Error executing feh command:", err)
				}
			}()
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := "Choose a wallpaper\n\n"

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

