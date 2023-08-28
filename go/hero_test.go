package main

import "testing"

func TestNewHero(t *testing.T) {
	h := NewHero("Test Hero,100,10,10")

	if h.Name != "Test Hero" {
		t.Errorf("got %v, want %v", h.Name, "Test Hero")
	}
	if h.Health != 100 {
		t.Errorf("got %v, want %v", h.Health, 100)
	}
	if h.AttackDamage != 10 {
		t.Errorf("got %v, want %v", h.AttackDamage, 10)
	}
	if h.Speed != 10 {
		t.Errorf("got %v, want %v", h.Speed, 10)
	}
}

func TestNewHeroPanics(t *testing.T) {
	tests := []struct {
		name string
		s    string
	}{
		{
			name: "Less attributes",
			s:    "Test Hero,100,10",
		},
		{
			name: "More attributes",
			s:    "Test Hero,100,10,10,10",
		},
		{
			name: "Separator is not comma",
			s:    "Test Hero|100|10|10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("The code did not panic")
				}
			}()

			NewHero(tt.s)
		})
	}
}

func TestHeroIsAlive(t *testing.T) {
	tests := []struct {
		name string
		h    Hero
		want bool
	}{
		{
			name: "healthy hero is alive",
			h: Hero{
				CharacterAttributes: &CharacterAttributes{
					Health: initialHealth,
				},
			},
			want: true,
		},
		{
			name: "injured hero is alive",
			h: Hero{
				CharacterAttributes: &CharacterAttributes{
					Health: minimalHealth,
				},
			},
			want: true,
		},
		{
			name: "dead hero is not alive",
			h: Hero{
				CharacterAttributes: &CharacterAttributes{
					Health: 0,
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.h.IsAlive()
			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeroHit(t *testing.T) {
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

	hero.Hit(&monster)

	if monster.Health != 90 {
		t.Errorf("got %v, want %v", monster.Health, 90)
	}
}

func TestHeroTake(t *testing.T) {
	monster := Monster{
		CharacterAttributes: &CharacterAttributes{
			Health:       initialHealth,
			AttackDamage: 10,
			Speed:        10,
			Name:         "Test Monster",
		},
		SpeedDamage: 10,
	}

	sword := Item{
		CharacterAttributes: CharacterAttributes{
			Name:         "Might Sword",
			Health:       0,
			AttackDamage: 50,
			Speed:        50,
		},
	}

	hero := Hero{
		CharacterAttributes: &CharacterAttributes{
			Health:       initialHealth,
			AttackDamage: 10,
			Speed:        10,
			Name:         "Test Hero",
		},
	}

	hero.Hit(&monster)
	if monster.Health != 90 {
		t.Errorf("Hitting without sword. got %v, want %v", monster.Health, 90)
	}

	hero.Take(sword)
	hero.Hit(&monster)

	if hero.AttackDamage != 60 {
		t.Errorf("got %v, want %v", hero.AttackDamage, 60)
	}
	if monster.Health != 30 {
		t.Errorf("Hitting with sword. got %v, want %v", monster.Health, 30)
	}
}
