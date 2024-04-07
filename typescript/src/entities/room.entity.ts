import { PlayerAttributes } from "../types";
import { Item, Monster } from "./";

export class Room {
  public name: string;
  public monsters: Monster[];
  private item: Item;

  constructor(
    name: string,
    monster: Monster,
    item: Item,
  ) {
    this.name = name;
    this.monsters = [monster];
    this.item = item;
  }

  public killMonster(monster: Monster) {
    const index = this.monsters.indexOf(monster);
    this.monsters.splice(index, 1);
  }

  /**
   * Initial guess is there will be max 2 monsters, so after each turn we revers order
   * Like that first monster should be on turn to fight
   */
  public changeMonstersFightOrder() {
    this.monsters = this.monsters.reverse();
  }

  public addMonster(clone: Monster) {
    this.monsters.push(clone);
  }

  public numberOfMonsters(): number {
    return this.monsters.length;
  }

  public monsterOnTurn() {
    return this.monsters[0];
  }

  public getItem(): PlayerAttributes {
    return this.item.getAttributes();
  }
};
