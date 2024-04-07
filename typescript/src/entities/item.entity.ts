import { Attributes } from "./basics.entity";
import { PlayerAttributes } from "../types";

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
