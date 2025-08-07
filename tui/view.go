package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/luketucich/dungeon-crawler/player"
)

type model struct {
	room   [][]rune
	player player.Player
}

func NewModel(room [][]rune, p player.Player) tea.Model {
	return model{
		room:   room,
		player: p,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "w":
			m.player.Move(0, -1, m.room)
		case "down", "s":
			m.player.Move(0, 1, m.room)
		case "left", "a":
			m.player.Move(-1, 0, m.room)
		case "right", "d":
			m.player.Move(1, 0, m.room)
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	var out strings.Builder

	for y := 0; y < len(m.room); y++ {
		for x := 0; x < len(m.room[y]); x++ {
			var cell string
			if m.player.X == x && m.player.Y == y {
				cell = lipgloss.NewStyle().
					Background(lipgloss.Color("#ff0000")). // bright red bg
					Bold(true).
					Render("  ")
			} else {
				tile := m.room[y][x]
				switch tile {
				case '#':
					cell = lipgloss.NewStyle().
						Background(lipgloss.Color("#555")).
						Render("  ")
				case '.':
					cell = lipgloss.NewStyle().
						Background(lipgloss.Color("#ccc")).
						Render("  ")
				default:
					cell = "  "
				}
			}

			out.WriteString(cell)
		}
		out.WriteString("\n")
	}

	return out.String()
}
