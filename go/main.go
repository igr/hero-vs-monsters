package main

import (
	"os"
	"path/filepath"
	"strings"
)

var tv Television = ConsoleTV{}

func main() {
	game := loadGame(filepath.Join("..", "game1.txt"))

	game.Play()
}

func loadGame(f string) Game {
	// load game from file
	content, err := os.ReadFile(f)
	if err != nil {
		panic(err)
	}

	// convert content into array of string lines
	lines := strings.Split(string(content), "\n")

	// first line is the hero
	hero := NewHero(lines[0])

	rooms := []Room{}
	for _, line := range lines[1:] {
		// skip empty lines
		if line == "" {
			continue
		}

		rooms = append(rooms, NewRoom(line))
	}

	return Game{
		hero: hero,
		maze: rooms,
	}
}
