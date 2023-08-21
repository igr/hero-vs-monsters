package hvm.app

import hvm.app.Event.*

/**
 * The main function that executes the game and returns the list of events that happened during the game.
 */
fun play(game: Game): List<Event> {
	val events = mutableListOf<Event>()

	var hero: Hero = game.hero
	var heroLastEvent: Event = HeroWins(hero)

	for (room in game.rooms) {
		hero = battleInRoom(hero, room).also { events.addAll(it.second) }.first

		if (hero.isDead()) {
			heroLastEvent = HeroDies(hero)
			break
		}

		hero = pickItem(hero, room.item).also { events += ItemPickEvent(it, room.item) }
	}

	return events + heroLastEvent
}

private fun pickItem(hero: Hero, item: Item): Hero {
	with(hero) {
		return Hero(
			name,
			health + item.health,
			attack + item.attack,
			speed + item.speed
		)
	}
}

private fun battleInRoom(heroAtRoomEntrance: Hero, room: Room): Pair<Hero, List<Event>> {
	val events = mutableListOf<Event>()
	events += RoomEvent(room, heroAtRoomEntrance)

	val aliveMonsters = mutableListOf(room.monster)

	var hero = heroAtRoomEntrance
	while (aliveMonsters.isNotEmpty()) {
		val monster = aliveMonsters.removeFirst()

		val fightEvent = fightHeroAndMonster(hero, monster).also { events += it }
		hero = fightEvent.hero

		if (hero.isDead()) {
			// we break the loop, but the event is added in the calling function, that is misleading
			break
		}

		val monsterAfterFight = fightEvent.monster
		if (monsterAfterFight.isDead()) {
			events += MonsterIsDeadEvent(monsterAfterFight)
			continue
		}

		cloneMonsterIfPossible(monsterAfterFight).run {
			aliveMonsters.addAll(this)
			if (this.size > 1) {
				events += MonsterClonedEvent(this.first())
			}
		}
	}

	// since the usage of Pair is immediate, we can use it as a return value
	return Pair(hero, events)
}

private fun fightHeroAndMonster(hero: Hero, monster: Monster) =
	if (monster.speed > hero.speed) {
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


private fun hitMonsterByHero(monster: Monster, hero: Hero) = monster.copy(health = monster.health - hero.attack)

private fun hitHeroByMonster(hero: Hero, monster: Monster) = hero.copy(health = hero.health - monster.attack)

private fun cloneMonsterIfPossible(monster: Monster): List<Monster> {
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

