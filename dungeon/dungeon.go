package dungeon

func CreateRoom(width, height int) [][]rune {
	room := make([][]rune, height)

	for y := 0; y < height; y++ {
		room[y] = make([]rune, width)

		for x := 0; x < width; x++ {
			if y == 0 || y == height-1 || x == 0 || x == width-1 {
				room[y][x] = '#' // Wall
			} else {
				room[y][x] = '.' // Floor
			}
		}
	}

	return room
}
