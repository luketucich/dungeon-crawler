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
	return model{room: room, player: p}
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyMsg); ok {
		switch key.String() {
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
	var b strings.Builder
	render := func(bg string, bold bool) string {
		style := lipgloss.NewStyle().Background(lipgloss.Color(bg))
		if bold {
			style = style.Bold(true)
		}
		return style.Render("  ")
	}

	for y := range m.room {
		for x := range m.room[y] {
			if m.player.X == x && m.player.Y == y {
				b.WriteString(render("#0096FF", true))
				continue
			}

			switch m.room[y][x] {
			case '#':
				b.WriteString(render("#555", false))
			case '.':
				b.WriteString(render("#ccc", false))
			default:
				b.WriteString("  ")
			}
		}
		b.WriteByte('\n')
	}

	return b.String()
}
