import { Item } from "./basics.entity";
import { Hero } from "./hero.entity";
import { Monster } from "./monster.entity";

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

  public takeOutMonster(monster: Monster) {
    // console.log('removing monster', monster.name);
    const index = this.monsters.indexOf(monster);
    this.monsters.splice(index, 1);
  }
};
