package hvm;

import java.util.List;
import java.util.Optional;

/**
 * Encapsulates the logic of cloning monsters.
 */
public class MonsterCloner {

	private final List<Monster> monsters;

	public MonsterCloner(List<Monster> monsters) {
		this.monsters = monsters;
	}

	public List<Monster> spawnClonedMonsters() {
		return monsters.stream()
				.filter(GameCharacter::isAlive)
				.map(Monster::spawnClone)
				.filter(Optional::isPresent)
				.map(Optional::get)
				.toList();

	}
}
