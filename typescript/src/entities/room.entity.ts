import { Item, Monster } from "./";

export class Room {
  public name: string;
  public monsters: Monster[];
  public item: Item;

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

  public pickItem(): Item {
    return this.item;
  }
};
