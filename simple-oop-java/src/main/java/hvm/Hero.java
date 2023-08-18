package hvm;

/**
 * Hero is a GameCharacter that can fight Monsters.
 */
public class Hero extends GameCharacter {

	public Hero(String name, int health, int attack, int speed) {
		super(name, health, attack, speed);
	}

	/**
	 * Fights a monster.
	 * How do we know if behaviour belongs in Hero or Monster or GameCharacter?
	 */
	public void fight(Monster monster) {
		if (monster.getSpeed() > getSpeed()) {
			monster.hit(this);
			if (isAlive()) {
				hit(monster);
			}
		} else {
			hit(monster);
			if (monster.isAlive()) {
				monster.hit(this);
			}
		}
	}

}
