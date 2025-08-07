package player

type Player struct {
	X int
	Y int
}

func NewPlayer(x, y int) Player {
	return Player{X: x, Y: y}
}

func (p *Player) Move(dx, dy int, room [][]rune) {
	newX := p.X + dx
	newY := p.Y + dy

	// Check bounds
	if newY < 0 || newY >= len(room) || newX < 0 || newX >= len(room[0]) {
		return
	}

	// Check collision with wall
	if room[newY][newX] == '#' {
		return
	}

	p.X = newX
	p.Y = newY
}
