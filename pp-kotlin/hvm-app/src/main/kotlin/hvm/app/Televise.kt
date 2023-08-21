package hvm.app

import tv.ShowOnTv

fun televiseGameEvents(events: List<Event>, showOnTv: ShowOnTv) {
	events.forEach {
		when (it) {
			is Event.RoomEvent -> showOnTv("Hero ${it.hero.name.value} enters ${it.room.name.value}")
			is Event.FightEvent -> {
				showOnTv("Monster ${it.monster.name.value} attacks: ${makeRoar()}")
				showOnTv("Hero ${it.hero.name.value} fights ${it.monster.name.value}")
			}
			is Event.MonsterIsDeadEvent -> showOnTv("Monster ${it.monster.name.value} is dead")
			is Event.ItemPickEvent -> showOnTv("Hero ${it.hero.name.value} founds ${it.item.name.value}")
			is Event.MonsterClonedEvent -> showOnTv("Monster ${it.monster.name.value} cloned!")
			is Event.HeroDies -> showOnTv("Hero ${it.hero.name.value} dies!")
			is Event.HeroWins -> showOnTv("Hero ${it.hero.name.value} wins!")
		}
	}
}