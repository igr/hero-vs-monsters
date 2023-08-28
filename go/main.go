package main

import (
	"flag"
	"os"
	"strings"
)

func main() {
	//var defaultGameFile = filepath.Join("..", "game1.txt")
	var tv Television = ConsoleTV{}

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

	game.tv = tv

	game.Play()
}
