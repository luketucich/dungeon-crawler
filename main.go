package main

import (
	"fmt"
	"github.com/luketucich/dungeon-crawler/dungeon"
)

func main() {
	room := dungeon.CreateRoom(6, 6)

	for _, row := range room {
		fmt.Println(string(row))
	}
}
