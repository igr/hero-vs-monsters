
import { Player } from './basics.entity';


export class Hero extends Player {
  constructor(attackDamage: number, health: number, speed: number, name: string) {
    super(attackDamage, health, speed, name);
  }

  public attack(monsterName: string): number {
    console.log(`Hero ${this.name} fights ${monsterName}`);
    return this.attackDamage;
  }

  public entersTheRoom(roomName: string): void {
    console.log(`Hero ${this.name} enters ${roomName} room!`);
  }

  public enhanceWithItem(attackDamage: number, health: number, speed: number, itemName: string): void {
    console.log(`Hero ${this.name} founds ${itemName}`);
    this.attackDamage += attackDamage;
    this.health += health;
    this.speed += speed;
  }

  public takeHitAndContinue(attackDamage: number, speedDamage: number): boolean {
    this.health -= attackDamage;
    this.speed -= speedDamage;
    return this.isAlive();
  };
}
