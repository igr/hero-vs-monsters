package main

import (
	"fmt"
	"os"
	"strings"
)

type Television interface {
	Show(message string)
}

type ConsoleTV struct{}

func (ConsoleTV) Show(message string) {
	fmt.Println(message)
}

type Game struct {
	tv   Television
	hero *Hero
	maze []Room
}

func (g Game) Play() {
	hero := g.hero
	for _, room := range g.maze {
		g.tv.Show("🚪 Hero " + hero.Name + " enters " + room.Name)

		aliveMonsters := room.AliveMonsters()
		for {
			for _, monster := range aliveMonsters {
				if hero.Speed > monster.Speed {
					g.tv.Show("🗡️ Hero " + hero.Name + " fights " + monster.Name)
					hero.Hit(monster)

					if monster.IsAlive() {
						g.tv.Show("🧌 Monster " + monster.Name + " attacks: " + monster.Roar())
						monster.Hit(hero)
					}
				} else {
					g.tv.Show("🧌 Monster " + monster.Name + " attacks: " + monster.Roar())
					monster.Hit(hero)
					if hero.IsAlive() {
						g.tv.Show("🗡️ Hero " + hero.Name + " fights " + monster.Name)
						hero.Hit(monster)
					}
				}

				if !monster.IsAlive() {
					g.tv.Show("💀 Monster " + monster.Name + " is dead")
					continue
				}

				if monster.CanBeCloned() {
					cloned := monster.clone()
					room.Monsters = append(room.Monsters, &cloned)
					g.tv.Show("👥 Monster " + monster.Name + " cloned itself!")
				}

				if !hero.IsAlive() {
					// do not process more rooms, as the hero died
					g.tv.Show("💀 Hero " + hero.Name + " died in room " + room.Name)
					return
				}
			}

			// refresh the list of alive monsters in case they cloned themselves
			aliveMonsters = room.AliveMonsters()
			if len(aliveMonsters) == 0 {
				g.tv.Show("✨ Hero " + hero.Name + " found " + room.Item.Name)
				hero.Take(room.Item)
				break
			}
		}
	}

	g.tv.Show("🏆 Hero " + hero.Name + " wins!")
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
