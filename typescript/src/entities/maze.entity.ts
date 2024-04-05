import { Hero, Monster, Room } from "./";

export class Maze {
  public rooms: Room[] = [];
  public hero: Hero;

  constructor(hero: Hero) {
    this.hero = hero;
  }

  public addRoom(room: Room): void {
    this.rooms.push(room);
  }

  private welcomeMessage(numberOfRooms: number): void {
    console.log("------------------------------------------------------------");
    console.log("Welcome to the monsters vs hero battle");
    console.log(`Maze has ${numberOfRooms} rooms with different monsters`);
    console.log("Let's start the fight!");
  }


  private heroAttacksMonster(monster: Monster): boolean | Monster {
    const attackDamage = this.hero.attack(monster.getName());
    return monster.isAliveOrCLonedAfterAttack(attackDamage);
  }

  private monsterAttacksHero(monster: Monster): boolean {
    const { attackDamage, speedDamage } = monster.roarAndAttack();
    return this.hero.isAliveAfterAttack(attackDamage, speedDamage, monster.getName());
  }


  public startFight(): void {
    this.welcomeMessage(this.rooms.length);

    this.rooms.forEach((room: Room) => {
      const isHeroFirst = this.hero.isHeroFirstAndGreeting(room.name, room.monsters[0].getSpeed());

      while (room.monsters.length) {
        if (isHeroFirst) {
          const isMonsterAliveOrCloned = this.heroAttacksMonster(room.monsters[0]);

          if (isMonsterAliveOrCloned === false) {
            room.killMonster(room.monsters[0]);
            if (!room.monsters.length) {
              this.hero.enhanceWithItem(room.item.getAttributes());
              continue;
            }
          } else {
            const isHeroStillAlive = this.monsterAttacksHero(room.monsters[0]);
            if (!isHeroStillAlive) break;

            if (room.monsters.length > 1) {
              room.monsters = room.monsters.reverse();
            } else if (isMonsterAliveOrCloned instanceof Monster) {
              room.monsters = [isMonsterAliveOrCloned, room.monsters[0]];
            }
          }
        } else {
          const isHeroStillAlive = this.monsterAttacksHero(room.monsters[0]);
          if (!isHeroStillAlive) break;

          const isMonsterAliveOrCloned = this.heroAttacksMonster(room.monsters[0]);

          if (isMonsterAliveOrCloned === false) {
            room.killMonster(room.monsters[0]);
            if (!room.monsters.length) {
              this.hero.enhanceWithItem(room.item.getAttributes());
              continue;
            }
          } else {
            if (room.monsters.length > 1) {
              room.monsters = room.monsters.reverse();
            } else if (isMonsterAliveOrCloned instanceof Monster) {
              room.monsters = [isMonsterAliveOrCloned, room.monsters[0]];
            }
          }
        }
      }
    });
  }
}
