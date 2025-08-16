package dungeon

import (
	"github.com/luketucich/dungeon-crawler/misc"
	"math/rand"
)

const (
	smallRoomArea  = 36
	mediumRoomArea = 100
	largeRoomArea  = 256
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

func CreateRoom(sizes ...int) Room {
	var width, height int

	if len(sizes) == 2 {
		width, height = sizes[0], sizes[1]
	} else {
		width, height = generateRoomSize()
	}

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
	addDoors(&room, generateDoorCount(width*height))
	return room
}

func addDoors(room *Room, count int) {
	if len(room.PossibleDoors) == 0 || count <= 0 {
		return
	}
	if count > len(room.PossibleDoors) {
		count = len(room.PossibleDoors)
	}

	misc.ShuffleSlice(room.PossibleDoors)

	for i := 0; i < count; i++ {
		x, y := room.PossibleDoors[i][0], room.PossibleDoors[i][1]
		room.Tiles[y][x].Structure = "door"
	}
}

func generateDoorCount(roomArea int) int {
	if roomArea <= smallRoomArea {
		return 1
	} else if roomArea <= mediumRoomArea {
		return 2
	} else {
		return 3
	}
}

func generateRoomSize() (width, height int) {
	roll := rand.Intn(100)
	switch {
	case roll < 50:
		return misc.Sqrt(smallRoomArea), misc.Sqrt(smallRoomArea)
	case roll < 85:
		return misc.Sqrt(mediumRoomArea), misc.Sqrt(mediumRoomArea)
	default:
		return misc.Sqrt(largeRoomArea), misc.Sqrt(largeRoomArea)
	}
}

func updateTile(room *Room, x, y int, structure string) {
	room.Tiles[y][x] = Tile{x, y, structure}
}
