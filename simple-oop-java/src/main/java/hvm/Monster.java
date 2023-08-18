package hvm;

public class Monster extends GameCharacter {

	private final int speedDamage;
	private boolean cloneable;

	public final int initialHealth;

	private final RoarFactory roarFactory = new RoarFactory();

	public Monster(String name, int health, int attack, int speed, int speedDamage, boolean cloneable) {
		super(name, health, attack, speed);
		this.initialHealth = health;
		this.speedDamage = speedDamage;
		this.cloneable = cloneable;
	}

	/**
	 * Monster hits the target and slows it down. This is a specialization of the orifinal behaviour.
	 */
	@Override
	public void hit(GameCharacter target) {
		super.hit(target);
		target.slowDown(speedDamage);
	}

	/**
	 * Spawns a clone of the monster if it is cloneable and its health is less than 25% of the initial health.
	 * If the monster is not cloned, returns {@code null}.
	 */
	public Monster spawnClone() {
		if (!cloneable) {
			return null;
		}
		if (health > this.initialHealth / 4) {
			return null;
		}
		var clonedMonster = new Monster(name + " clone", health / 2, attack, speed, speedDamage, false);
		this.health = health / 2;
		this.cloneable = false;
		return clonedMonster;
	}

	public String roar() {
		return roarFactory.roar();
	}
}