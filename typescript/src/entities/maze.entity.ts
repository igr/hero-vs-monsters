import { Item } from "./basics.entity";
import { Hero } from "./hero.entity";
import { Monster } from "./monster.entity";
import { Room } from "./room.entity";

export class Maze {
  public rooms: Room[] = [];
  public hero: Hero;

  constructor(hero: Hero) {
    this.hero = hero;
  }


  public createHero(attackDamage: number, health: number, speed: number, name: string): void {
    this.hero = new Hero(attackDamage, health, speed, name);
  }

  public createMonster(attackDamage: number, health: number, speed: number, speedDamage: number, name: string, cloneable: boolean): Monster {
    return new Monster(attackDamage, health, speed, speedDamage, name, cloneable);
  }

  public createItem(name: string, attackDamage: number, health: number, speed: number): Item {
    return new Item(name, attackDamage, health, speed);
  }

  public createRoom(name: string, monster: Monster, item: Item,): void {
    const room = new Room(name, monster, item);
    this.rooms.push(room);
  }

  public startFight(): void {
    console.log("------------------------------------------------------------");
    console.log("Welcome to the monsters vs hero battle");
    console.log(`Maze has ${this.rooms.length} rooms with different monsters`);
    console.log("Let's start the fight!");
    this.rooms.forEach((room: Room) => {
      console.log("------------------------------------------------------------");
      this.hero.entersTheRoom(room.name);
      const isHeroFirst = true;
      // const isHeroFirst = this.hero.speed > room.monsters[0].speed;

      while (room.monsters.length) {
        // console.log('🚀 ~ Maze ~ this.rooms.forEach ~ room.monsters:', room.monsters);
        // console.log('🚀 ~ Maze ~ this.rooms.forEach ~ hero:', this.hero);
        if (isHeroFirst) {
          const attackDamage = this.hero.attack(room.monsters[0].name);
          const doWeContinueOrClonedMonster = room.monsters[0].takeHitAndContinue(attackDamage);

          if (doWeContinueOrClonedMonster instanceof Monster) {
            console.log('🚀 ~ Maze ~ this.rooms.forEach ~ doWeContinueOrClonedMonster:', doWeContinueOrClonedMonster);
            room.monsters.push(doWeContinueOrClonedMonster);
            const { attackDamage, speedDamage } = room.monsters[0].roarAndAttack();
            const doWeContinue = this.hero.takeHitAndContinue(attackDamage, speedDamage);
            continue;
          }

          if (doWeContinueOrClonedMonster) {
            const { attackDamage, speedDamage } = room.monsters[0].roarAndAttack();
            const doWeContinue = this.hero.takeHitAndContinue(attackDamage, speedDamage);
            if (!doWeContinue) {
              console.log('Hero defeated');
              break;
            }
          } else {
            const { attackDamage, health, speed } = room.item.getAttributes();
            this.hero.enhanceWithItem(attackDamage, health, speed, room.item.name);
            room.takeOutMonster(room.monsters[0]);
          }
        } {

        }
      }

      // const firstHit = heroSpeed > this.room.monster.speed
    });
  }

  // start fight
  // next room
  // pick first hit 
};
