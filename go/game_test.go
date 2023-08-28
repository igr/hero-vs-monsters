package main

import (
	"path/filepath"
	"strings"
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

type testTV struct {
	lines []string
}

func (t *testTV) Show(s string) {
	t.lines = append(t.lines, s)
}

func TestPlayOneSingleRoomSlowerHeroHeroWins(t *testing.T) {
	hero := NewHero("Beorn,240,10,20")
	r := NewRoom("Hallway,Haunthand,20,6,40,8,false,Sword,0,4,5")

	// hero will take 2 turns to take down the monster
	// monster will take 120 turns in take down the hero
	turnsToWin := 2

	linesTV := &testTV{}

	g := Game{
		hero: hero,
		maze: []Room{r},
		tv:   linesTV,
	}

	g.Play()

	lines := linesTV.lines

	if len(lines) != 8 {
		t.Errorf("got %v, want %v", len(lines), 8)
	}

	if !strings.EqualFold(lines[0], "🚪 Hero Beorn enters Hallway") {
		t.Errorf("got %v, want %v", lines[0], "🚪 Hero Beorn enters Hallway")
	}

	firstTurn := 1
	// first turn: monster is faster
	if !strings.HasPrefix(lines[firstTurn], "🧌 Monster Haunthand attacks:") {
		t.Errorf("got %v, want %v", lines[firstTurn], "🧌 Monster Haunthand attacks:")
	}
	if !strings.EqualFold(lines[firstTurn+1], "🗡️ Hero Beorn fights Haunthand") {
		t.Errorf("got %v, want %v", lines[firstTurn+1], "🗡️ Hero Beorn fights Haunthand")
	}
	// last turn: monster is still faster
	lastTurn := turnsToWin*2 - firstTurn // 3
	if !strings.HasPrefix(lines[lastTurn], "🧌 Monster Haunthand attacks:") {
		t.Errorf("got %v, want %v", lines[lastTurn], "🧌 Monster Haunthand attacks:")
	}
	if !strings.EqualFold(lines[lastTurn+1], "🗡️ Hero Beorn fights Haunthand") {
		t.Errorf("got %v, want %v", lines[lastTurn+1], "🗡️ Hero Beorn fights Haunthand")
	}
	if !strings.EqualFold(lines[lastTurn+2], "💀 Monster Haunthand is dead") {
		t.Errorf("got %v, want %v", lines[lastTurn+2], "💀 Monster Haunthand is dead")
	}
	if !strings.EqualFold(lines[lastTurn+3], "✨ Hero Beorn found Sword") {
		t.Errorf("got %v, want %v", lines[lastTurn+3], "✨ Hero Beorn found Sword")
	}
	if !strings.EqualFold(lines[lastTurn+4], "🏆 Hero Beorn wins!") {
		t.Errorf("got %v, want %v", lines[lastTurn+4], "🏆 Hero Beorn wins!")
	}
}

func TestPlayOneSingleRoomFasterHeroHeroWins(t *testing.T) {
	hero := NewHero("Beorn,240,10,60")
	r := NewRoom("Hallway,Haunthand,20,6,40,8,false,Sword,0,4,5")

	// hero will take 2 turns to take down the monster
	// monster will take 120 turns in take down the hero
	turnsToWin := 2

	linesTV := &testTV{}

	g := Game{
		hero: hero,
		maze: []Room{r},
		tv:   linesTV,
	}

	g.Play()

	lines := linesTV.lines

	if len(lines) != 7 {
		t.Errorf("got %v, want %v", len(lines), 7)
	}

	if !strings.EqualFold(lines[0], "🚪 Hero Beorn enters Hallway") {
		t.Errorf("got %v, want %v", lines[0], "🚪 Hero Beorn enters Hallway")
	}

	firstTurn := 1
	// first turn: hero is faster
	if !strings.EqualFold(lines[firstTurn], "🗡️ Hero Beorn fights Haunthand") {
		t.Errorf("got %v, want %v", lines[firstTurn], "🗡️ Hero Beorn fights Haunthand")
	}
	if !strings.HasPrefix(lines[firstTurn+1], "🧌 Monster Haunthand attacks:") {
		t.Errorf("got %v, want %v", lines[firstTurn+2], "🧌 Monster Haunthand attacks:")
	}
	// last turn: hero is still faster
	lastTurn := turnsToWin*2 - firstTurn // 3
	if !strings.EqualFold(lines[lastTurn], "🗡️ Hero Beorn fights Haunthand") {
		t.Errorf("got %v, want %v", lines[lastTurn], "🗡️ Hero Beorn fights Haunthand")
	}
	if !strings.EqualFold(lines[lastTurn+1], "💀 Monster Haunthand is dead") {
		t.Errorf("got %v, want %v", lines[lastTurn+1], "💀 Monster Haunthand is dead")
	}
	// hero wins
	if !strings.EqualFold(lines[lastTurn+2], "✨ Hero Beorn found Sword") {
		t.Errorf("got %v, want %v", lines[lastTurn+2], "✨ Hero Beorn found Sword")
	}
	if !strings.EqualFold(lines[lastTurn+3], "🏆 Hero Beorn wins!") {
		t.Errorf("got %v, want %v", lines[lastTurn+3], "🏆 Hero Beorn wins!")
	}
}

func TestPlayOneSingleRoomSlowerHeroMonsterWins(t *testing.T) {
	hero := NewHero("Beorn,240,10,20")
	r := NewRoom("Tower,Red Dragon,250,40,200,100,true,Gold,0,0,0")

	// monster will take 6 turns in take down the hero
	// hero will take 13 turns to take down the monster
	turnsToWin := 6

	linesTV := &testTV{}

	g := Game{
		hero: hero,
		maze: []Room{r},
		tv:   linesTV,
	}

	g.Play()

	lines := linesTV.lines

	if len(lines) != 13 {
		t.Errorf("got %v, want %v", len(lines), 13)
	}

	if !strings.EqualFold(lines[0], "🚪 Hero Beorn enters Tower") {
		t.Errorf("got %v, want %v", lines[0], "🚪 Hero Beorn enters Tower")
	}
	// first turn: monster is faster
	firstTurn := 1
	if !strings.HasPrefix(lines[firstTurn], "🧌 Monster Red Dragon attacks:") {
		t.Errorf("got %v, want %v", lines[firstTurn], "🧌 Monster Red Dragon attacks:")
	}
	if !strings.EqualFold(lines[firstTurn+1], "🗡️ Hero Beorn fights Red Dragon") {
		t.Errorf("got %v, want %v", lines[firstTurn+1], "🗡️ Hero Beorn fights Red Dragon")
	}
	// last turn: monster is faster
	lastTurn := turnsToWin*2 - firstTurn // 11
	if !strings.HasPrefix(lines[lastTurn], "🧌 Monster Red Dragon attacks:") {
		t.Errorf("got %v, want %v", lines[lastTurn], "🧌 Monster Red Dragon attacks:")
	}
	// hero dies
	if !strings.EqualFold(lines[lastTurn+1], "💀 Hero Beorn died in room Tower") {
		t.Errorf("got %v, want %v", lines[lastTurn+1], "💀 Hero Beorn died in room Tower")
	}
}

func TestPlayOneSingleRoomFasterHeroMonsterWins(t *testing.T) {
	hero := NewHero("Beorn,240,10,600")
	r := NewRoom("Tower,Red Dragon,250,40,200,100,true,Gold,0,0,0")

	// monster will take 6 turns in take down the hero
	// hero will take 13 turns to take down the monster
	turnsToWin := 6

	linesTV := &testTV{}

	g := Game{
		hero: hero,
		maze: []Room{r},
		tv:   linesTV,
	}

	g.Play()

	lines := linesTV.lines

	if len(lines) != 13 {
		t.Errorf("got %v, want %v", len(lines), 13)
	}

	if !strings.EqualFold(lines[0], "🚪 Hero Beorn enters Tower") {
		t.Errorf("got %v, want %v", lines[0], "🚪 Hero Beorn enters Tower")
	}
	// first turn: hero is faster
	firstTurn := 1
	if !strings.EqualFold(lines[firstTurn], "🗡️ Hero Beorn fights Red Dragon") {
		t.Errorf("got %v, want %v", lines[firstTurn], "🗡️ Hero Beorn fights Red Dragon")
	}
	if !strings.HasPrefix(lines[firstTurn+1], "🧌 Monster Red Dragon attacks:") {
		t.Errorf("got %v, want %v", lines[firstTurn+1], "🧌 Monster Red Dragon attacks:")
	}
	// last turn: hero is no longer faster
	lastTurn := turnsToWin*2 - firstTurn // 11
	if !strings.HasPrefix(lines[lastTurn], "🧌 Monster Red Dragon attacks:") {
		t.Errorf("got %v, want %v", lines[lastTurn], "🧌 Monster Red Dragon attacks:")
	}
	// hero dies
	if !strings.EqualFold(lines[lastTurn+1], "💀 Hero Beorn died in room Tower") {
		t.Errorf("got %v, want %v", lines[lastTurn+1], "💀 Hero Beorn died in room Tower")
	}
}

func TestPlayOneSingleRoomClonableMonsterHeroWins(t *testing.T) {
	hero := NewHero("Beorn,300,10,20")
	r := NewRoom("Tower,Red Dragon,50,40,200,100,true,Gold,0,0,0")

	// monster will take 6 turns in take down the hero
	// hero will take 5 turns to take down one monster
	// but the monster will clone at turn 4
	// hero will take both in two more turns
	turnsToWin := 5 + 1

	linesTV := &testTV{}

	g := Game{
		hero: hero,
		maze: []Room{r},
		tv:   linesTV,
	}

	g.Play()

	lines := linesTV.lines

	if len(lines) != 18 {
		t.Errorf("got %v, want %v", len(lines), 18)
	}

	if !strings.EqualFold(lines[0], "🚪 Hero Beorn enters Tower") {
		t.Errorf("got %v, want %v", lines[0], "🚪 Hero Beorn enters Tower")
	}
	// first turn: monster is faster
	firstTurn := 1
	if !strings.HasPrefix(lines[firstTurn], "🧌 Monster Red Dragon attacks:") {
		t.Errorf("got %v, want %v", lines[firstTurn], "🧌 Monster Red Dragon attacks:")
	}
	if !strings.EqualFold(lines[firstTurn+1], "🗡️ Hero Beorn fights Red Dragon") {
		t.Errorf("got %v, want %v", lines[firstTurn+1], "🗡️ Hero Beorn fights Red Dragon")
	}
	// last turn: monsters are faster
	lastTurn := turnsToWin*2 - firstTurn + 1 + 1 // one entry for the clone, one for the death of the first monster
	if !strings.EqualFold(lines[lastTurn-1], "💀 Monster Red Dragon is dead") {
		t.Errorf("got %v, want %v", lines[lastTurn-1], "💀 Monster Red Dragon is dead")
	}
	if !strings.HasPrefix(lines[lastTurn], "🧌 Monster Red Dragon (Cloned) attacks:") {
		t.Errorf("got %v, want %v", lines[lastTurn], "🧌 Monster Red Dragon (Cloned) attacks:")
	}
	if !strings.EqualFold(lines[lastTurn+1], "🗡️ Hero Beorn fights Red Dragon (Cloned)") {
		t.Errorf("got %v, want %v", lines[lastTurn+1], "🗡️ Hero Beorn fights Red Dragon (Cloned)")
	}
	if !strings.EqualFold(lines[lastTurn+2], "💀 Monster Red Dragon (Cloned) is dead") {
		t.Errorf("got %v, want %v", lines[lastTurn+2], "💀 Monster Red Dragon (Cloned) is dead")
	}
	// hero wins
	if !strings.EqualFold(lines[lastTurn+3], "✨ Hero Beorn found Gold") {
		t.Errorf("got %v, want %v", lines[lastTurn+3], "✨ Hero Beorn found Gold")
	}
	if !strings.EqualFold(lines[lastTurn+4], "🏆 Hero Beorn wins!") {
		t.Errorf("got %v, want %v", lines[lastTurn+4], "🏆 Hero Beorn wins!")
	}
}

func TestPlayOneSingleRoomClonableMonsterHeroDies(t *testing.T) {
	hero := NewHero("Beorn,240,10,20")
	r := NewRoom("Tower,Red Dragon,60,40,200,100,true,Gold,0,0,0")

	// monster will take 6 turns in take down the hero
	// hero will take 6 turns to take down one monster
	// but the monster will clone at turn 5
	// monsters will take hero in one more turn
	turnsToWin := 5 + 1

	linesTV := &testTV{}

	g := Game{
		hero: hero,
		maze: []Room{r},
		tv:   linesTV,
	}

	g.Play()

	lines := linesTV.lines

	if len(lines) != 14 {
		t.Errorf("got %v, want %v", len(lines), 14)
	}

	if !strings.EqualFold(lines[0], "🚪 Hero Beorn enters Tower") {
		t.Errorf("got %v, want %v", lines[0], "🚪 Hero Beorn enters Tower")
	}
	// first turn: monster is faster
	firstTurn := 1
	if !strings.HasPrefix(lines[firstTurn], "🧌 Monster Red Dragon attacks:") {
		t.Errorf("got %v, want %v", lines[firstTurn], "🧌 Monster Red Dragon attacks:")
	}
	if !strings.EqualFold(lines[firstTurn+1], "🗡️ Hero Beorn fights Red Dragon") {
		t.Errorf("got %v, want %v", lines[firstTurn+1], "🗡️ Hero Beorn fights Red Dragon")
	}
	// last turn: monsters are faster
	lastTurn := turnsToWin*2 - firstTurn + 1 // one entry for the clone
	if !strings.HasPrefix(lines[lastTurn], "🧌 Monster Red Dragon attacks:") {
		t.Errorf("got %v, want %v", lines[lastTurn], "🧌 Monster Red Dragon attacks:")
	}
	// hero dies
	if !strings.EqualFold(lines[lastTurn+1], "💀 Hero Beorn died in room Tower") {
		t.Errorf("got %v, want %v", lines[lastTurn+1], "💀 Hero Beorn died in room Tower")
	}
}
