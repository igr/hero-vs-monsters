package main

import (
	"math/rand"
)

// ClonabilityThreshold is the percentage of health below which a monster can be cloned
const ClonabilityThreshold = 25

type Monster struct {
	*CharacterAttributes
	SpeedDamage   int
	Clonable      bool
	Cloned        bool
	initialHealth int
}

func (m *Monster) CanBeCloned() bool {
	return m.IsAlive() && m.Clonable && !m.Cloned && (m.Health <= m.initialHealth*ClonabilityThreshold/100)
}

// clone returns a new Monster with half of its healt.
// It will also mark the current monster as cloned, so it cannot be cloned again.
func (m *Monster) clone() Monster {
	newHealth := m.Health / 2

	cloned := Monster{
		CharacterAttributes: &CharacterAttributes{
			Name:         m.Name + " (Cloned)",
			Health:       newHealth,
			AttackDamage: m.AttackDamage,
			Speed:        m.Speed,
		},
		SpeedDamage:   m.SpeedDamage,
		Clonable:      false,
		Cloned:        false,
		initialHealth: newHealth,
	}

	m.Health -= newHealth
	m.Cloned = true // important to avoid infinite cloning

	return cloned
}

func (m *Monster) Roar() string {
	single := []rune{'H', 'W', 'L'}
	multiple := []rune{'R', 'O', 'A'}

	roar := ""

	// repeat 10 times the roar generation
	for i := 0; i < 10; i++ {
		base := single
		baseCount := 1
		random := rand.Intn(100)
		if random%2 == 0 {
			base = multiple
			randomCount := rand.Intn(3)
			baseCount = 3 + randomCount
		}

		// from the base, generate baseCount runes
		// and convert them into a string

		for i := 0; i < baseCount; i++ {
			random := rand.Intn(len(base))
			roar += string(base[random])
		}
	}

	return roar
}

func (m *Monster) IsAlive() bool {
	return m.Health > 0
}

func (m *Monster) Hit(h *Hero) {
	if !m.IsAlive() {
		return
	}

	h.Health -= m.AttackDamage
	// the monster slows down the hero
	h.Speed -= m.SpeedDamage
}
