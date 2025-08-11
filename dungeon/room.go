package dungeon

import (
	"math/rand"
	"time"
)

type Room struct {
	Tiles         [][]Tile
	PossibleDoors [][]int // each: [x, y]
	Width, Height int
}

type Tile struct {
	x, y      int
	Structure string
}

func CreateTile(x, y int, structure string) Tile {
	return Tile{x, y, structure}
}

func isCorner(x, y, width, height int) bool {
	return (x == 0 || x == width-1) && (y == 0 || y == height-1)
}

func isBorder(x, y, width, height int) bool {
	return y == 0 || y == height-1 || x == 0 || x == width-1
}

func shuffleSlice[T any](s []T) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
}

func CreateRoom(width, height int) Room {
	tiles := make([][]Tile, height)
	perimeterCap := max(0, (width-2)*2+(height-2)*2)
	possibleDoors := make([][]int, 0, perimeterCap)

	for y := 0; y < height; y++ {
		tiles[y] = make([]Tile, width)
		for x := 0; x < width; x++ {
			if isBorder(x, y, width, height) {
				tiles[y][x] = CreateTile(x, y, "wall")
				if !isCorner(x, y, width, height) {
					possibleDoors = append(possibleDoors, []int{x, y})
				}
			} else {
				tiles[y][x] = CreateTile(x, y, "floor")
			}
		}
	}

	room := Room{tiles, possibleDoors, width, height}
	addDoors(room, generateDoorCount(room))

	return room
}

func addDoors(room Room, count int) {
	if len(room.PossibleDoors) == 0 || count <= 0 {
		return
	}
	if count > len(room.PossibleDoors) {
		count = len(room.PossibleDoors)
	}

	shuffleSlice(room.PossibleDoors)

	for i := 0; i < count; i++ {
		x, y := room.PossibleDoors[i][0], room.PossibleDoors[i][1]
		room.Tiles[y][x].Structure = "door"
	}
}

func generateDoorCount(room Room) int {
	roomArea := room.Width * room.Height

	if roomArea <= 16 {
		return 1
	} else if roomArea <= 81 {
		return 2
	} else if roomArea <= 160 {
		return 3
	} else {
		return 4
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
