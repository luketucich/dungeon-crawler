package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/luketucich/dungeon-crawler/dungeon"
	"github.com/luketucich/dungeon-crawler/player"
)

type model struct {
	room   dungeon.Room
	player player.Player
}

func NewModel(r dungeon.Room, p player.Player) tea.Model {
	return model{room: r, player: p}
}

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

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
			return m, tea.Batch(tea.ExitAltScreen, tea.Quit)
		}
	}
	return m, nil
}

const (
	tileW       = 2
	playerGlyph = "@ "
)

var (
	wallStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#4A4A4A"))
	floorStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("#6E6E6E"))
	doorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#A9A9A9")).Bold(true).Underline(true)
	playerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#E4C44A")).Bold(true)
	borderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#3A3A3A"))
	hudStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#9C9C9C")).Faint(true)
)

const (
	wallGlyph = "▒▒"
	floorA    = "· "
	floorB    = " ."
	doorGlyph = "╬╬"
	voidGlyph = "  "
)

func (m model) drawTile(x, y int) string {
	if m.player.X == x && m.player.Y == y {
		return playerStyle.Render(playerGlyph)
	}

	switch m.room.Tiles[y][x].Structure {
	case "wall":
		return wallStyle.Render(wallGlyph)
	case "floor":
		if (x+y)%2 == 0 {
			return floorStyle.Render(floorA)
		}
		return floorStyle.Render(floorB)
	case "door":
		return doorStyle.Render(doorGlyph)
	default:
		return voidGlyph
	}
}

func (m model) hud() string {
	return hudStyle.Render("WASD to move • Q to quit  |  @ = You   ╬╬ = Door   ▒ = Wall   · = Floor")
}

func (m model) View() string {
	var b strings.Builder

	b.WriteString(borderStyle.Render("┌" + strings.Repeat("─", m.room.Width*tileW) + "┐"))
	b.WriteByte('\n')

	for y := 0; y < m.room.Height; y++ {
		b.WriteString(borderStyle.Render("│"))
		for x := 0; x < m.room.Width; x++ {
			b.WriteString(m.drawTile(x, y))
		}
		b.WriteString(borderStyle.Render("│"))
		b.WriteByte('\n')
	}

	b.WriteString(borderStyle.Render("└" + strings.Repeat("─", m.room.Width*tileW) + "┘"))
	b.WriteString("\n\n")
	b.WriteString(m.hud())

	return b.String()
}
