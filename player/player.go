package player

import (
	"github.com/luketucich/dungeon-crawler/dungeon"
	"github.com/luketucich/dungeon-crawler/misc"
)

type Player struct {
	X         int
	Y         int
	inventory []misc.Item
}

func NewPlayer(x, y int, inventory []misc.Item) Player {
	return Player{x, y, inventory}
}

func (p *Player) Move(dx, dy int, room dungeon.Room) {
	newX, newY := p.X+dx, p.Y+dy
	if newX < 0 || newY < 0 || newY >= room.Height || newX >= room.Width {
		return
	}

	newTile := room.Tiles[newY][newX]
	if newTile.Structure == "wall" {
		return
	}

	p.X, p.Y = newX, newY
}

func (p *Player) GrabItem(item misc.Item) {
	p.inventory = append(p.inventory, item)
}

func (p *Player) DropItem(indexToRemove int) {
	p.inventory = append(p.inventory[:indexToRemove], p.inventory[indexToRemove+1:]...)
}
