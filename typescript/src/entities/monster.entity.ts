
import { Player } from "./basics.entity";
import { MonsterAttributes } from "./types";

export class Monster extends Player {
  private cloneableHealth: number;
  private speedDamage: number;
  private cloneable: boolean;
  public isClone: boolean = false;

  constructor(input: MonsterAttributes) {
    super(input);
    this.speedDamage = input.speedDamage;
    this.cloneableHealth = 0.25 * this.health;
    this.cloneable = input.cloneable;
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
      const repeatCount: number = letterRepeatMap[randomElement]();
      roar = roar + randomElement.repeat(repeatCount);
      letters = letters.filter(item => item !== randomElement);
    }

    console.log(`Monster ${this.name} attacks: ${roar}`);

    return { attackDamage: this.attackDamage, speedDamage: this.speedDamage, };
  }

  public isAliveOrCLonedAfterAttack(attackDamage: number): boolean | Monster {
    this.health -= attackDamage;

    if (this.cloneable && this.health < this.cloneableHealth && this.health > 0) {
      return this.cloneMonster();
    }
    // console.log("monster take a hit, is it alive", this.isAlive());
    return this.isAlive();
  };

  public cloneMonster(): Monster {
    console.log("Cloning monster!", this.name);
    if (!this.cloneable) throw Error('Monster could not be cloned!');
    this.health = 0.5 * this.health;

    const cloned = new Monster({
      attackDamage: this.attackDamage,
      health: this.health,
      speed: this.speed,
      name: `${this.name} clone`,
      speedDamage: this.speedDamage,
      cloneable: false
    });
    cloned.isClone = true;
    console.log("Monster cloned!", cloned.name, cloned.cloneableHealth, cloned.isClone);
    return cloned;
  }
};



