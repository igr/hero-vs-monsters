package hvm;
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

	public boolean isAlive() {
		return health > 0;
	}

	public void hit(GameCharacter target) {
		target.health -= attack;
	}

	public void slowDown(int speedDamage) {
		speed -= speedDamage;
	}
	public void apply(Item item) {
		health += item.health();
		attack += item.attack();
		speed += item.speed();
	}
}
