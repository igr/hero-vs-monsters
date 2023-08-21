package hvm;

import hvm.tv.Television;

import java.util.List;

/**
 * The main game logic.
 */
public final class Game {
	private final Hero hero;
	private final List<Room> rooms;

	private final Television tv;

	public Game(Hero hero, List<Room> rooms, Television tv) {
		this.hero = hero;
		this.rooms = rooms;
		this.tv = tv;
	}

	/**
	 * The main game loop.
	 */
	public void play() {
		// we need to iterate over rooms
		for (Room room : rooms) {
			tv.show("Hero " + hero.getName() + " enters " + room.name());

			final var monsters = room.monsters();

			// iterate over all alive monsters in the room
			while (true) {
				if (hero.isDead()) {
					return;
				}

				var allMonstersAlive = monsters
						.stream()
						.map(Monster::isAlive)
						.reduce(Boolean::logicalOr)
						.orElse(false);

				if (!allMonstersAlive) {
					break;
				}

				monsters.stream()
						.filter(Monster::isAlive)
						.forEach(this::fightMonster);

				final var clonedMonsters = new MonsterCloner(monsters).spawnClonedMonsters();
				if (!clonedMonsters.isEmpty()) {
					tv.show("Monster " + monsters.get(0).getName() + " cloned!");
					monsters.addAll(clonedMonsters);
				}
			}

			final var item = room.item();
			tv.show("Hero " + hero.getName() + " founds " + item.name());

			hero.useItem(item);
		}

		tv.show("Hero " + hero.getName() + " wins!");
	}

	/**
	 * Fighting monsters.
	 * How to be sure that method belongs here? Should we move it to Hero class?
	 */
	private void fightMonster(final Monster monster) {
		if (hero.isDead()) {
			return;
		}
		final var roar = monster.roar();
		tv.show("Monster " + monster.getName() + " attacks: " + roar);

		tv.show("Hero " + hero.getName() + " fights " + monster.getName());
		hero.fight(monster);

		if (hero.isDead()) {
			tv.show("Hero " + hero.getName() + " is dead");
		}
		if (monster.isDead()) {
			tv.show("Monster " + monster.getName() + " is dead");
		}
	}

}
