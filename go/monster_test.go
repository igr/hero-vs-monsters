package main

import "testing"

const initialHealth = 100
const minimalHealth = 25

func TestCanBeCloned(t *testing.T) {
	tests := []struct {
		name string
		m    Monster
		want bool
	}{
		{
			name: "healthy clonable monster cannot be cloned",
			m: Monster{
				CharacterAttributes: &CharacterAttributes{
					Health: initialHealth,
				},
				Clonable:      true,
				initialHealth: initialHealth,
			},
			want: false,
		},
		{
			name: "injured clonable monster can be cloned",
			m: Monster{
				CharacterAttributes: &CharacterAttributes{
					Health: minimalHealth,
				},
				Clonable:      true,
				initialHealth: initialHealth,
			},
			want: true,
		},
		{
			name: "healthy non-clonable cannot be cloned",
			m: Monster{
				CharacterAttributes: &CharacterAttributes{
					Health: initialHealth,
				},
				Clonable:      false,
				initialHealth: initialHealth,
			},
			want: false,
		},
		{
			name: "injured non-clonable cannot be cloned",
			m: Monster{
				CharacterAttributes: &CharacterAttributes{
					Health: minimalHealth,
				},
				Clonable:      false,
				initialHealth: initialHealth,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.CanBeCloned()
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClone(t *testing.T) {
	monster := Monster{
		CharacterAttributes: &CharacterAttributes{
			Name:   "Test",
			Health: 81, // using an odd number to force a non-integer division
		},
		Clonable: true,
	}

	// force a clone withouht the guard clause
	cloned := monster.clone()

	if cloned.Health != 40 {
		t.Errorf("got %v, want %v", cloned.Health, 40)
	}
	if cloned.Name != monster.Name+" (Cloned)" {
		t.Errorf("got %v, want %v", cloned.Name, monster.Name+" (Cloned)")
	}
	if cloned.AttackDamage != monster.AttackDamage {
		t.Errorf("got %v, want %v", cloned.AttackDamage, monster.AttackDamage)
	}
	if cloned.Speed != monster.Speed {
		t.Errorf("got %v, want %v", cloned.Speed, monster.Speed)
	}
	if cloned.SpeedDamage != monster.SpeedDamage {
		t.Errorf("got %v, want %v", cloned.SpeedDamage, monster.SpeedDamage)
	}
	if cloned.initialHealth != 40 {
		t.Errorf("got %v, want %v", cloned.initialHealth, 40)
	}

	// assert values for the original monster
	if monster.Health != 41 {
		t.Errorf("got %v, want %v", monster.Health, 41)
	}
	if monster.Cloned != true {
		t.Errorf("got %v, want %v", monster.Cloned, true)
	}
}

func TestMonsterIsAlive(t *testing.T) {
	tests := []struct {
		name string
		m    Monster
		want bool
	}{
		{
			name: "healthy monster is alive",
			m: Monster{
				CharacterAttributes: &CharacterAttributes{
					Health: initialHealth,
				},
			},
			want: true,
		},
		{
			name: "injured monster is alive",
			m: Monster{
				CharacterAttributes: &CharacterAttributes{
					Health: minimalHealth,
				},
			},
			want: true,
		},
		{
			name: "dead monster is not alive",
			m: Monster{
				CharacterAttributes: &CharacterAttributes{
					Health: 0,
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.IsAlive()
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMonsterHit(t *testing.T) {
	monster := Monster{
		CharacterAttributes: &CharacterAttributes{
			Health:       initialHealth,
			AttackDamage: 10,
			Speed:        10,
			Name:         "Test Monster",
		},
		SpeedDamage: 10,
	}

	hero := Hero{
		CharacterAttributes: &CharacterAttributes{
			Health:       initialHealth,
			AttackDamage: 10,
			Speed:        10,
			Name:         "Test Hero",
		},
	}

	monster.Hit(&hero)

	if hero.Health != 90 {
		t.Errorf("got %v, want %v", hero.Health, 90)
	}
	if hero.Speed != 0 {
		t.Errorf("got %v, want %v", hero.Speed, 0)
	}
}
