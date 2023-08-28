package main

import "testing"

func TestAliveMonsters(t *testing.T) {
	r := NewRoom("Hallway,Haunthand,20,6,40,8,false,Sword,0,4,5")

	alive := r.AliveMonsters()
	if len(alive) != 1 {
		t.Errorf("got %v, want %v", len(alive), 1)
	}

	// kill the only monster
	r.Monsters[0].Health = 0

	alive = r.AliveMonsters()
	if len(alive) != 0 {
		t.Errorf("got %v, want %v", len(alive), 0)
	}

	// add 10 alive monster
	for i := 0; i < 10; i++ {
		r.Monsters = append(r.Monsters, &Monster{
			CharacterAttributes: &CharacterAttributes{
				Health: initialHealth,
			},
		})
	}
	alive = r.AliveMonsters()
	if len(alive) != 10 {
		t.Errorf("got %v, want %v", len(alive), 10)
	}

	// kill two random monsters
	r.Monsters[3].Health = 0
	r.Monsters[7].Health = 0

	alive = r.AliveMonsters()
	if len(alive) != 8 {
		t.Errorf("got %v, want %v", len(alive), 8)
	}
}

func TestNewRoom(t *testing.T) {
	r := NewRoom("Hallway,Haunthand,20,6,40,8,false,Sword,0,4,5")

	if r.Name != "Hallway" {
		t.Errorf("got %v, want %v", r.Name, "Hallway")
	}
	if len(r.Monsters) != 1 {
		t.Errorf("got %v, want %v", len(r.Monsters), 1)
	}

	// eval monster
	monster := r.Monsters[0]
	if monster.Name != "Haunthand" {
		t.Errorf("got %v, want %v", monster.Name, "Haunthand")
	}
	if monster.Health != 20 {
		t.Errorf("got %v, want %v", monster.Health, 20)
	}
	if monster.initialHealth != 20 {
		t.Errorf("got %v, want %v", monster.initialHealth, 20)
	}
	if monster.AttackDamage != 6 {
		t.Errorf("got %v, want %v", monster.AttackDamage, 6)
	}
	if monster.Speed != 40 {
		t.Errorf("got %v, want %v", monster.Speed, 40)
	}
	if monster.SpeedDamage != 8 {
		t.Errorf("got %v, want %v", monster.SpeedDamage, 8)
	}
	if monster.Clonable != false {
		t.Errorf("got %v, want %v", monster.Clonable, false)
	}
	if monster.Cloned != false {
		t.Errorf("got %v, want %v", monster.Clonable, false)
	}

	// eval item in the room
	item := r.Item
	if item.Name != "Sword" {
		t.Errorf("got %v, want %v", item.Name, "Sword")
	}
	if item.Health != 0 {
		t.Errorf("got %v, want %v", item.Health, 0)
	}
	if item.AttackDamage != 4 {
		t.Errorf("got %v, want %v", item.AttackDamage, 4)
	}
	if item.Speed != 5 {
		t.Errorf("got %v, want %v", item.Speed, 5)
	}
}

func TestNewRoomPanics(t *testing.T) {
	tests := []struct {
		name string
		s    string
	}{
		{
			name: "Less attributes",
			s:    "Test Room,100,10",
		},
		{
			name: "More attributes",
			s:    "Dark Room,Helltree,30,10,5,24,false,Shield,5,0,10,73",
		},
		{
			name: "Separator is not comma",
			s:    "Dark Room|Helltree|30|10|5|24|false|Shield|5|0|10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("The code did not panic")
				}
			}()

			NewRoom(tt.s)
		})
	}
}
