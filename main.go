package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/luketucich/dungeon-crawler/dungeon"
	"github.com/luketucich/dungeon-crawler/player"
	"github.com/luketucich/dungeon-crawler/tui"
)

func main() {
	room := dungeon.CreateRoom(10, 6)
	p := player.NewPlayer(4, 4) // Starting position (not a wall)
	prog := tea.NewProgram(tui.NewModel(room, p))
	if _, err := prog.Run(); err != nil {
		log.Fatal(err)
	}
}
