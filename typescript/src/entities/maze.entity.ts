import { Hero, Monster, Room } from "./";
import { logger } from "../logger";


export class Maze {
  private rooms: Room[] = [];
  private hero: Hero;

  constructor(hero: Hero) {
    this.hero = hero;
  }

  public addRoom(room: Room): void {
    this.rooms.push(room);
  }

  private welcomeMessage(numberOfRooms: number): void {
    logger("------------------------------------------------------------");
    logger("Welcome to the monsters vs hero battle");
    logger(`Maze has ${numberOfRooms} rooms with different monsters`);
    logger("Let's start the fight!");
  }

  /**
   * @param monster 
   * @returns Cloned monster instance or boolean for isAlive status 
   */
  private heroAttacksMonster(monster: Monster): boolean | Monster {
    const attackDamage = this.hero.attack(monster.getName());
    return monster.isAliveOrCLonedAfterAttack(attackDamage);
  }

  private monsterAttacksHero(monster: Monster): boolean {
    const { attackDamage, speedDamage } = monster.roarAndAttack();
    return this.hero.isAliveAfterAttack(attackDamage, speedDamage);
  }


  public startFight(): void {
    this.welcomeMessage(this.rooms.length);

    for (const room of this.rooms) {
      this.hero.entersRoomMessage(room.name);
      const isHeroFirst = this.hero.isHeroFirst(room.monsterOnTurn().getSpeed());

      while (room.numberOfMonsters() > 0) {
        if (isHeroFirst) {
          //** Hero first hits
          const isMonsterAliveOrCloned = this.heroAttacksMonster(room.monsterOnTurn());
          if (isMonsterAliveOrCloned instanceof Monster) room.addMonster(isMonsterAliveOrCloned);

          if (isMonsterAliveOrCloned === false) {
            room.killMonster(room.monsterOnTurn());
            if (room.numberOfMonsters() === 0) {
              this.hero.pickUpItem(room.getItem());
              continue;
            }
          } else {
            const isHeroStillAlive = this.monsterAttacksHero(room.monsterOnTurn());
            if (!isHeroStillAlive) {
              room.monsterOnTurn().winsTheBattle();
              break;
            }
          }

          room.changeMonstersFightOrder();
        } else {
          //** Monster first hits
          const isHeroStillAlive = this.monsterAttacksHero(room.monsterOnTurn());
          if (!isHeroStillAlive) {
            room.monsterOnTurn().winsTheBattle();
            break;
          }
          const isMonsterAliveOrCloned = this.heroAttacksMonster(room.monsterOnTurn());
          if (isMonsterAliveOrCloned instanceof Monster) room.addMonster(isMonsterAliveOrCloned);

          if (isMonsterAliveOrCloned === false) {
            room.killMonster(room.monsterOnTurn());
            if (!room.monsters.length) {
              this.hero.pickUpItem(room.getItem());
              continue;
            }
          }
          room.changeMonstersFightOrder();
        }
      }
    }

    if (this.hero.isAlive()) this.hero.winsTheBattle();
  }
}
