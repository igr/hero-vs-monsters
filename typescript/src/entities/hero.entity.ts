
import { Player } from './basics.entity';
import { PlayerAttributes } from '../types';
import { logger } from "../logger";

export class Hero extends Player {
  constructor(input: PlayerAttributes) {
    super(input);
  }

  public attack(monsterName: string): number {
    logger(`Hero ${this.name} fights ${monsterName}`);
    return this.attackDamage;
  }

  public isHeroFirst(monsterSpeed: number): boolean {
    return this.speed > monsterSpeed;
  }

  public entersRoomMessage(roomName: string) {
    logger("------------------------------------------------------------");
    logger(`Hero ${this.name} enters ${roomName} room!`);
    logger("------------------------------------------------------------");
  }

  public pickUpItem(item: PlayerAttributes): void {
    logger(`Hero ${this.name} wins the battle`);
    logger(`Hero ${this.name} founds ${item.name}`);
    this.attackDamage += item.attackDamage;
    this.health += item.health;
    this.speed += item.speed;
  }

  public isAliveAfterAttack(attackDamage: number, speedDamage: number): boolean {
    this.health -= attackDamage;
    this.speed -= speedDamage;

    if (this.isAlive()) return true;

    this.playerDiedMessage();
    return false;
  };
}
