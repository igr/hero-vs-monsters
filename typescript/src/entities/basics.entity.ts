import { PlayerAttributes } from "../types";
import { logger } from "../logger";

export abstract class Attributes {
  protected attackDamage: number;
  protected health: number;
  protected speed: number;
  protected name: string;

  constructor(input: PlayerAttributes) {
    this.name = input.name;
    this.attackDamage = input.attackDamage;
    this.health = input.health;
    this.speed = input.speed;
  }

  public getName(): string {
    return this.name;
  }

  public getSpeed(): number {
    return this.speed;
  }
}

export abstract class Player extends Attributes {
  constructor(input: PlayerAttributes) {
    super(input);
  }

  public isAlive() {
    return this.health <= 0 ? false : true;
  }

  public playerDiedMessage(): void {
    logger(`${this.constructor.name} ${this.name} is dead!`);
  }

  public winsTheBattle(): void {
    logger(`🚀🚀🚀 ${this.constructor.name} ${this.name} wins! 🚀🚀🚀`);
  }
}
