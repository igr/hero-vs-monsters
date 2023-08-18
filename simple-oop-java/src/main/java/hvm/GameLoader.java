package hvm;

import hvm.tv.ConsoleTv;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;

public class GameLoader {

	public Game loadGame() throws IOException {
		final var gameLines = Files.readAllLines(Path.of("../game1.txt"));
		final var hero = parseHero(gameLines.remove(0));
		final var rooms = gameLines.stream().map(this::parseRoom).toList();
		return new Game(hero, rooms, new ConsoleTv());
	}

	private Hero parseHero(String line) {
		final var heroParts = line.split(",");
		final var name = heroParts[0].trim();
		final var health = Integer.parseInt(heroParts[1]);
		final var attack = Integer.parseInt(heroParts[2]);
		final var speed = Integer.parseInt(heroParts[3]);
		return new Hero(name, health, attack, speed);
	}

	private Room parseRoom(String line) {
		final var roomParts = line.split(",");
		final var roomName = roomParts[0].trim();
		final var monster = parseMonster(roomParts);
		final var item = parseItem(roomParts);
		return new Room(roomName, monster, item);
	}

	private Monster parseMonster(String[] lines) {
		final var name = lines[1].trim();
		final var health = Integer.parseInt(lines[2]);
		final var attack = Integer.parseInt(lines[3]);
		final var speed = Integer.parseInt(lines[4]);
		final var speedDamage = Integer.parseInt(lines[5]);
		final var duplicable = Boolean.parseBoolean(lines[6]);
		return new Monster(name, health, attack, speed, speedDamage, duplicable);
	}

	private Item parseItem(String[] lines) {
		final var name = lines[7].trim();
		final var health = Integer.parseInt(lines[8]);
		final var attack = Integer.parseInt(lines[9]);
		final var speed = Integer.parseInt(lines[10]);
		return new Item(name, health, attack, speed);
	}
}
