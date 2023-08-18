package hvm;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import java.util.Random;
import java.util.function.Function;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

public class Monster extends GameCharacter {

	private int speedDamage;

	public Monster(String name, int health, int attack, int speed, int speedDamage) {
		super(name, health, attack, speed);
		this.speedDamage = speedDamage;
	}

	@Override
	public void hit(GameCharacter target) {
		super.hit(target);
		target.slowDown(speedDamage);
	}

	private static final List<String> roar = "HWLROA".chars().mapToObj(Character::toString).toList();

	public String roar() {
		var rnd = new Random();

		var list = new ArrayList<>(roar);
		Collections.shuffle(list);

		return list.stream().map(c -> {
			switch (c) {
				case "H", "W", "L" -> {
					return c;
				}
				case "R", "O", "A" -> {
					return String.join("", Collections.nCopies(rnd.nextInt(3, 7), c));
				}
				default -> {
					return "";
				}
			}
		}).collect(Collectors.joining());
	}
}