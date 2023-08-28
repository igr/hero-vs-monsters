package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Fighter interface {
	IsAlive() bool
	Hit(b Fighter)
}

type CharacterAttributes struct {
	AttackDamage int
	Health       int
	Name         string
	Speed        int
}

func (c *CharacterAttributes) String() string {
	return fmt.Sprintf("%s  (Health: %d, Attack damage: %d, Speed: %d)", c.Name, c.Health, c.AttackDamage, c.Speed)
}

type Room struct {
	Name     string
	Monsters []Monster
	Item     Item
}

// NewRoom parses a string and returns a Room
func NewRoom(s string) Room {
	tokens := strings.Split(s, ",")
	if len(tokens) != 11 {
		panic("invalid room")
	}

	monsterHealth := parseIntAttribute(tokens[2])

	monster := Monster{
		CharacterAttributes: &CharacterAttributes{
			Name:         tokens[1],
			Health:       monsterHealth,
			AttackDamage: parseIntAttribute(tokens[3]),
			Speed:        parseIntAttribute(tokens[4]),
		},
		SpeedDamage:   parseIntAttribute(tokens[5]),
		Clonable:      parseBoolAttribute(tokens[6]),
		initialHealth: monsterHealth,
	}

	item := Item{
		CharacterAttributes: CharacterAttributes{
			Name:         tokens[7],
			Health:       parseIntAttribute(tokens[8]),
			AttackDamage: parseIntAttribute(tokens[9]),
			Speed:        parseIntAttribute(tokens[10]),
		},
	}

	return Room{
		Name:     tokens[0],
		Monsters: []Monster{monster},
		Item:     item,
	}
}

func (r *Room) AliveMonsters() []Monster {
	monsters := []Monster{}
	for _, m := range r.Monsters {
		if m.IsAlive() {
			monsters = append(monsters, m)
		}
	}

	return monsters
}

func (r *Room) Combat(h *Hero, m *Monster) {
	for {
		if !h.IsAlive() {
			tv.Show("💀 Hero " + h.Name + " dies!")
			break
		}
		if !m.IsAlive() {
			tv.Show("💀 Monster " + m.Name + " is dead")
			break
		}

		if h.Speed > m.Speed {
			h.Hit(m)

			if m.IsAlive() {
				tv.Show("🧌 Monster " + m.Name + " attacks: " + m.Roar())
				m.Hit(h)
			}
		} else {
			m.Hit(h)
			if h.IsAlive() {
				tv.Show("🧌 Monster " + m.Name + " attacks: " + m.Roar())
				h.Hit(m)
			}
		}

		if m.CanBeCloned() {
			cloned := m.clone()
			r.Monsters = append(r.Monsters, cloned)
			tv.Show("👥 Monster " + m.Name + " cloned itself!")
		}
	}
}

type Item struct {
	CharacterAttributes
}

// parsing an interger attribute will raise a panic if the string cannot be converted to an integer
func parseIntAttribute(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

// parsing a boolean attribute will raise a panic if the string cannot be converted to a boolean
func parseBoolAttribute(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		panic(err)
	}

	return b
}
