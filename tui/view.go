package tui

import (
	"github.com/luketucich/dungeon-crawler/dungeon"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/luketucich/dungeon-crawler/player"
)

type model struct {
	room   dungeon.Room
	player player.Player
}

func NewModel(r dungeon.Room, p player.Player) tea.Model {
	return model{room: r, player: p}
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

	for y := range m.room.Height {
		for x := range m.room.Width {
			if m.player.X == x && m.player.Y == y {
				b.WriteString(render("#0096FF", true))
				continue
			}

			switch m.room.Tiles[y][x].Structure {
			case "wall":
				b.WriteString(render("#555", false))
			case "floor":
				b.WriteString(render("#ccc", false))
			case "door":
				b.WriteString(render("#ccc", false))
			default:
				b.WriteString("  ")
			}
		}
		b.WriteByte('\n')
	}

	return b.String()
}
