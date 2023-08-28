package main

import (
	"fmt"
	"strconv"
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
