package hvm;

public final class Room {
	private final String name;
	private final Monster monster;
	private final Item item;

	public Room(String name, Monster monster, Item item) {
		this.name = name;
		this.monster = monster;
		this.item = item;
	}

	public String getName() {
		return name;
	}

	public Monster getMonster() {
		return monster;
	}

	public Item getItem() {
		return item;
	}
}
