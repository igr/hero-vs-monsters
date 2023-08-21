package hvm.app

@JvmInline
value class CharacterName(val value: String)

@JvmInline
value class ItemName(val value: String)

@JvmInline
value class Health(internal val value: Int) {
	operator fun minus(attack: Attack): Health {
		return Health(value - attack.value)
	}

	operator fun plus(health: Health): Health {
		return Health(this.value + health.value)
	}

	operator fun div(i: Int): Health {
		return Health(this.value / i)
	}
}

@JvmInline
value class Attack(internal val value: Int) {
	operator fun plus(attack: Attack): Attack {
		return Attack(this.value + attack.value)
	}
}

@JvmInline
value class Speed(private val value: Int) {
	operator fun compareTo(speed: Speed): Int {
		return this.value.compareTo(speed.value)
	}

	operator fun plus(speed: Speed): Speed {
		return Speed(this.value + speed.value)
	}
}

@JvmInline
value class SpeedDamage(private val value: Int)

@JvmInline
value class Cloneable(val value: Boolean)

/**
 * The only abstraction between Hero and Monster, more for esthetically reasons.
 */
interface LivingCreature {
	val health: Health
	fun isAlive() = health.value > 0
	fun isDead() = health.value <= 0
}

data class Hero(
	val name: CharacterName,
	override val health: Health,
	val attack: Attack,
	val speed: Speed) : LivingCreature {
}

data class Monster(
	val name: CharacterName,
	override val health: Health,
	val initialHealth: Health,
	val attack: Attack,
	val speed: Speed,
	val speedDamage: SpeedDamage,
	val cloneable: Cloneable
) : LivingCreature

data class Item(
	val name: ItemName,
	val health: Health,
	val attack: Attack,
	val speed: Speed,
)


@JvmInline
value class RoomName(val value: String)

data class Room(val name: RoomName, val monster: Monster, val item: Item)

data class Game(val hero: Hero, val rooms: List<Room>)
