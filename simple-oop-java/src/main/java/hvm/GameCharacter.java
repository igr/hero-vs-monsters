package hvm;

/**
 * As {@link Hero} and {@link Monster} have a lot in common, we can extract a common superclass.
 */
public abstract class GameCharacter {

	protected final String name;
	protected int health;
	protected int attack;
	protected int speed;

	public GameCharacter(String name, int health, int attack, int speed) {
		this.name = name;
		this.health = health;
		this.attack = attack;
		this.speed = speed;
	}

	// This is a non-record style getter method, old-school style.
	public String getName() {
		return name;
	}

	public int getHealth() {
		return health;
	}

	public int getAttack() {
		return attack;
	}

	public int getSpeed() {
		return speed;
	}

	/**
	 * Returns {@code true} if the character is alive, {@code false} otherwise.
	 */
	public boolean isAlive() {
		return health > 0;
	}
	public boolean isDead() {
		return health <= 0;
	}

	/**
	 * Attacks the given target by reducing its health by the attack value.
	 */
	public void hit(GameCharacter target) {
		target.health -= attack;
	}

	/**
	 * Reduces the speed of the character by the given amount.
	 * Notice the different direction of {@link #hit(GameCharacter)} method and this method.
	 * The first one performs action on another character, while this one performs action on the character itself.
	 */
	public void slowDown(int speedDamage) {
		speed -= speedDamage;
	}

	public void useItem(Item item) {
		health += item.health();
		attack += item.attack();
		speed += item.speed();
	}
}
