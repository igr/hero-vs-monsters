package main

import "strings"

type Hero struct {
	*CharacterAttributes
	Items []Item
}

func (h *Hero) IsAlive() bool {
	return h.Health > 0
}

func NewHero(s string) *Hero {
	tokens := strings.Split(s, ",")
	if len(tokens) != 4 {
		panic("invalid hero")
	}

	return &Hero{
		CharacterAttributes: &CharacterAttributes{
			Name:         tokens[0],
			Health:       parseIntAttribute(tokens[1]),
			AttackDamage: parseIntAttribute(tokens[2]),
			Speed:        parseIntAttribute(tokens[3]),
		},
	}
}

func (h *Hero) Hit(m *Monster) {
	m.Health -= h.AttackDamage
}

func (h *Hero) Take(i Item) {
	h.Items = append(h.Items, i)

	h.Health += i.Health
	h.AttackDamage += i.AttackDamage
	h.Speed += i.Speed
}
