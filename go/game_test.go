package main

import (
	"path/filepath"
	"testing"
)

func TestLoadGame(t *testing.T) {
	g := loadGame(filepath.Join("testdata", "game1.txt"))

	hero := g.hero
	if hero.Name != "Beorn" {
		t.Errorf("got %v, want %v", hero.Name, "Beorn")
	}

	rooms := g.maze
	if len(rooms) != 3 {
		t.Errorf("got %v, want %v", len(rooms), 3)
	}
}
