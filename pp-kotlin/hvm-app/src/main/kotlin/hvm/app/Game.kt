package hvm.app

import hvm.app.Event.*

fun play(game: Game): List<Event> {
	var hero: Hero = game.hero
	val roomEvents = mutableListOf<Event>()
	var heroEvent: Event = HeroWins(hero)
	for (room in game.rooms) {
		val pair = battleInRoom(hero, room).also { roomEvents.addAll(it.second) }
		hero = pair.first
		if (hero.isDead()) {
			heroEvent = HeroDies(hero)
			break
		}
		hero = pickItem(hero, room.item)
		roomEvents.add(ItemPickEvent(hero, room.item))
	}
	return roomEvents + heroEvent
}

fun pickItem(hero: Hero, item: Item): Hero {
	with(hero) {
		return Hero(
			name,
			health + item.health,
			attack + item.attack,
			speed + item.speed
		)
	}
}

fun battleInRoom(freshHero: Hero, room: Room): Pair<Hero, List<Event>> {
	val fightEvents = mutableListOf<Event>()
	fightEvents.add(RoomEvent(room, freshHero))

	val aliveMonsters = mutableListOf(room.monster)

	var hero = freshHero
	while (aliveMonsters.isNotEmpty()) {
		val monster = aliveMonsters.removeFirst()

		val fightResult = fightHeroAndMonster(hero, monster).also { fightEvents.add(it) }
		hero = fightResult.hero
		if (hero.isDead()) {
			break
		}
		with(fightResult.monster) {
			if (this.isAlive()) {
				cloneMonsterIfPossible(this).run {
					aliveMonsters.addAll(this)
					if (this.size > 1) {
						fightEvents.add(MonsterClonedEvent(this.first()))
					}
				}
			} else {
				fightEvents.add(MonsterIsDeadEvent(this))
			}
		}
	}
	return Pair(hero, fightEvents)
}

private fun fightHeroAndMonster(hero: Hero, monster: Monster): FightEvent {
	return if (monster.speed > hero.speed) {
		val newHero = hitHeroByMonster(hero, monster)
		val newMonster = if (newHero.isAlive()) {
			hitMonsterByHero(monster, newHero)
		} else monster
		FightEvent(newHero, newMonster)
	} else {
		val newMonster = hitMonsterByHero(monster, hero)
		val newHero = if (newMonster.isAlive()) {
			hitHeroByMonster(hero, newMonster)
		} else hero
		FightEvent(newHero, newMonster)
	}
}


private fun hitMonsterByHero(monster: Monster, hero: Hero) = monster.copy(health = monster.health - hero.attack)

private fun hitHeroByMonster(hero: Hero, monster: Monster) = hero.copy(health = hero.health - monster.attack)
fun cloneMonsterIfPossible(monster: Monster): List<Monster> {
	if (!monster.cloneable.value) {
		return listOf(monster)
	}
	if (monster.health.value > monster.initialHealth.value / 4) {
		return listOf(monster)
	}
	return listOf(
		monster.copy(health = monster.health / 2),
		monster.copy(name = CharacterName("${monster.name.value} clone"), health = monster.health / 2)
	)
}

