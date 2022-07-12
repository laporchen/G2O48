package main;

import (
	"github.com/laporchen/go2048/internal/game"
	tea "github.com/charmbracelet/bubbletea"
	"fmt"
	"os"
)


func main() {
	p := tea.NewProgram(game.InitialModel())
    if err := p.Start(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}


