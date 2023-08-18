package hvm;

import hvm.tv.ConsoleTv;
import hvm.tv.Television;

import java.util.List;

public final class Game {
	private final Hero hero;
	private final List<Room> rooms;

	private final Television tv;

	public Game(Hero hero, List<Room> rooms, Television tv) {
		this.hero = hero;
		this.rooms = rooms;
		this.tv = tv;
	}
	public void play() {
		for (Room room : rooms) {
			tv.show("Hero " + hero.getName() + " enters " + room.getName());

			final var monster = room.getMonster();
			Monster monsterDuplicate = null;

			while (monster.isAlive()) {

				final var roar = monster.roar();
				tv.show("Monster " + monster.getName() + " roars " + roar);

				tv.show("Hero " + hero.getName() + " fights " + monster.getName());
				hero.fight(monster);

				if (!hero.isAlive()) {
					tv.show("Hero " + hero.getName() + " is dead");
					return;
				}

				if (!monster.isAlive()) {
					tv.show("Monster " + monster.getName() + " is dead");

					final var item = room.getItem();
					tv.show("Hero " + hero.getName() + " founds " + item.name());

					hero.apply(item);
				}
			}
		}

		tv.show("Hero " + hero.getName() + " wins!");
	}

}
