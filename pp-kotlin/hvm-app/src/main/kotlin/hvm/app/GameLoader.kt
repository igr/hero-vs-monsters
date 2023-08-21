package hvm.app

import java.nio.file.Files
import java.nio.file.Path

fun loadGame(): Game {
	val gameLines = Files.readAllLines(Path.of("../game1.txt"))
	val hero = parseHero(gameLines.removeAt(0))
	val rooms = gameLines.map { parseRoom(it) }
	return Game(hero, rooms)
}

private fun parseHero(line: String) =
	line.split(",").let {
		Hero(
			CharacterName(it[0].trim()),
			Health(it[1].toInt()),
			Attack(it[2].toInt()),
			Speed(it[3].toInt())
		)
	}

private fun parseRoom(line: String) =
	line.split(",").let {
		Room(RoomName(it[0].trim()), parseMonster(it), parseItem(it))
	}

private fun parseMonster(lines: List<String>) = Monster(
	CharacterName(lines[1].trim()),
	Health(lines[2].toInt()),
	Health(lines[2].toInt()),
	Attack(lines[3].toInt()),
	Speed(lines[4].toInt()),
	SpeedDamage(lines[5].toInt()),
	Cloneable(lines[6].toBoolean())
)

private fun parseItem(lines: List<String>) = Item(
	ItemName(lines[7].trim()),
	Health(lines[8].toInt()),
	Attack(lines[9].toInt()),
	Speed(lines[10].toInt())
)