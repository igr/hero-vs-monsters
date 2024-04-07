
import { Player } from "./basics.entity";
import { MonsterAttributes } from "../types";
import { logger } from "../logger";


export class Monster extends Player {
  private cloneableHealth: number;
  private speedDamage: number;

  constructor(input: MonsterAttributes) {
    super(input);
    this.speedDamage = input.speedDamage;
    this.cloneableHealth = input.cloneable ? 0.25 * this.health : 0;
  }

  public roarAndAttack() {
    let roar: string = '';

    function repeatBetween3and6Times(): number {
      return Math.floor(Math.random() * (6 - 3 + 1) + 3);
    }

    function repeatOnce(): number {
      return 1;
    }

    const letterRepeatMap: { [key: string]: () => number; } = {
      'H': repeatOnce,
      'W': repeatOnce,
      'l': repeatOnce,
      'R': repeatBetween3and6Times,
      'O': repeatBetween3and6Times,
      'A': repeatBetween3and6Times
    };

    let letters = ['H', 'W', 'l', 'R', 'O', 'A'];

    while (letters.length) {
      const randomElement = letters[Math.floor(Math.random() * letters.length)];
      const repeatCount = letterRepeatMap[randomElement]();
      roar = roar + randomElement.repeat(repeatCount);
      letters = letters.filter(item => item !== randomElement);
    }

    logger(`Monster ${this.name} attacks: ${roar}`);

    return { attackDamage: this.attackDamage, speedDamage: this.speedDamage, };
  }

  public isAliveOrCLonedAfterAttack(attackDamage: number): boolean | Monster {
    this.health -= attackDamage;

    if (this.health > 0 && this.health < this.cloneableHealth) {
      return this.cloneMonster();
    }

    if (this.isAlive()) return true;

    this.playerDiedMessage();
    return false;
  };

  public cloneMonster(): Monster {
    logger(`Cloning monster!, ${this.name}`);
    this.health = 0.5 * this.health;

    return new Monster({
      attackDamage: this.attackDamage,
      health: this.health,
      speed: this.speed,
      name: `${this.name} clone`,
      speedDamage: this.speedDamage,
      cloneable: false
    });
  }
};



