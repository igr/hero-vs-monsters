
import { Player } from './basics.entity';
import { PlayerAttributes } from './types';


export class Hero extends Player {
  constructor(input: PlayerAttributes) {
    super(input);
  }

  public attack(monsterName: string): number {
    console.log(`Hero ${this.name} fights ${monsterName}`);
    return this.attackDamage;
  }

  public isHeroFirst(monsterSpeed: number): boolean {
    return this.speed > monsterSpeed;
  }

  public isHeroFirstAndGreeting(roomName: string, monsterSpeed: number): boolean {
    console.log("------------------------------------------------------------");
    console.log(`Hero ${this.name} enters ${roomName} room!`);
    return this.isHeroFirst(monsterSpeed);
  }

  public enhanceWithItem(input: PlayerAttributes): void {
    console.log(`Hero ${this.name} founds ${input.name}`);
    this.attackDamage += input.attackDamage;
    this.health += input.health;
    this.speed += input.speed;
  }

  private heroLostMessage(monsterName: string): void {
    console.log("------------------------------------------------------------");
    console.log(`Hero ${this.name} lost the match! => Monster ${monsterName} wins!`);
    console.log("------------------------------------------------------------");
  }

  public isAliveAfterAttack(attackDamage: number, speedDamage: number, monsterName: string): boolean {
    this.health -= attackDamage;
    this.speed -= speedDamage;
    const isAlive = this.isAlive();
    if (!isAlive) this.heroLostMessage(monsterName);
    return isAlive;
  };
}
