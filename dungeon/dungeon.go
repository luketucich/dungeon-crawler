package dungeon

type Room struct {
	Tiles         [][]Tile
	Width, Height int
}

type Tile struct {
	x, y      int
	Structure string
}

func CreateTile(x, y int, structure string) Tile {
	return Tile{x, y, structure}
}

func CreateRoom(width, height int) Room {
	tiles := make([][]Tile, height)

	for y := 0; y < height; y++ {
		tiles[y] = make([]Tile, width)

		for x := 0; x < width; x++ {
			if y == 0 || y == height-1 || x == 0 || x == width-1 {
				tiles[y][x] = CreateTile(x, y, "wall")
			} else {
				tiles[y][x] = CreateTile(x, y, "floor")
			}
		}
	}

	return Room{tiles, width, height}
}
