package main

import "fmt"

type Television interface {
	Show(message string)
}

type ConsoleTV struct{}

func (ConsoleTV) Show(message string) {
	fmt.Println(message)
}

type Game struct {
	hero *Hero
	maze []Room
}

func (g Game) Play() {
	hero := g.hero
	for _, room := range g.maze {
		tv.Show("🚪 Hero " + hero.Name + " enters " + room.Name)

		for {
			aliveMonsters := room.AliveMonsters()
			if len(aliveMonsters) == 0 {
				break
			}

			for _, monster := range aliveMonsters {
				room.Combat(hero, &monster)
			}

			if !hero.IsAlive() {
				// do not process more rooms, as the hero died
				return
			}
		}
	}

	tv.Show("🏆 Hero " + hero.Name + " wins!")
}
