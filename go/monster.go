package main

import "math/rand"

type Monster struct {
	*CharacterAttributes
	SpeedDamage int
	Clonable    bool
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

	tv.Show("🧌 Monster " + m.Name + " attacks: " + m.Roar())

	h.Health -= m.AttackDamage
	// the monster slows down the hero
	h.Speed -= m.SpeedDamage
}
