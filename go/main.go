package main

import (
	"flag"
	"os"
	"strings"
)

var tv Television = ConsoleTV{}

func main() {
	//var defaultGameFile = filepath.Join("..", "game1.txt")

	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		tv.Show("invalid number of arguments")
		os.Exit(1)
	}

	gameFile := args[0]
	if !strings.HasSuffix(gameFile, ".txt") {
		tv.Show("invalid file name")
		os.Exit(1)
	}

	// check if file exists
	if _, err := os.Stat(gameFile); os.IsNotExist(err) {
		tv.Show("file does not exist")
		os.Exit(1)
	}

	game := loadGame(gameFile)

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
