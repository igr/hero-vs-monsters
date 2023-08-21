package hvm.app

sealed class Event {
	data class HeroWins(val hero: Hero) : Event()
	data class HeroDies(val hero: Hero) : Event()
	data class MonsterClonedEvent(val monster: Monster) : Event()
	data class ItemPickEvent(val hero: Hero, val item: Item) : Event()
	data class MonsterIsDeadEvent(val monster: Monster) : Event()
	data class RoomEvent(val room: Room, val hero: Hero) : Event()
	data class FightEvent(val hero: Hero, val monster: Monster) : Event()
}