import { PlayerAttributes } from "./types";

export class Attributes {
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

export class Item extends Attributes {
  constructor(input: PlayerAttributes) {
    super(input);
  }

  public getAttributes(): PlayerAttributes {
    return {
      attackDamage: this.attackDamage,
      health: this.health,
      speed: this.speed,
      name: this.name
    };
  }
}

export class Player extends Attributes {
  constructor(input: PlayerAttributes) {
    super(input);
  }

  public isAlive() {
    return this.health <= 0 ? false : true;
  }

  public winsTheBattle(): void {
    console.log(`${this.constructor.name} ${this.name} wins!`);
  }
}
