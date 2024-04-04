export class BasicAttributes {
  protected attackDamage: number;
  protected health: number;
  public speed: number;

  constructor(attackDamage: number, health: number, speed: number) {
    this.attackDamage = attackDamage;
    this.health = health;
    this.speed = speed;
  }
}

export class Item extends BasicAttributes {
  public name: string;
  constructor(name: string, attackDamage: number, health: number, speed: number) {
    super(attackDamage, health, speed);
    this.name = name;
  }

  public getAttributes() {
    return {
      attackDamage: this.attackDamage,
      health: this.health,
      speed: this.speed,
    };
  }
}

export class Player extends BasicAttributes {
  public name: string;

  constructor(attackDamage: number, health: number, speed: number, name: string) {
    super(attackDamage, health, speed);
    this.name = name;
  }

  isAlive() {
    if (this.health <= 0) {
      // console.log(`${this.constructor.name} ${this.name} is dead!`);
      return false;
    }
    return true;

  }

  public winsTheBattle(): void {
    console.log(`${this.constructor.name} ${this.name} wins!`);
  }
}
