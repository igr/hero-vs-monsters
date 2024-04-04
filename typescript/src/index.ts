import { Hero } from './entities/hero.entity';
import { Maze } from './entities/maze.entity';
import { readFileSync } from 'node:fs';
import { join } from 'node:path';

console.log("🚀🚀🚀 Initializing hero-vs-monsters in Node.js + Typescript");

const testFilePath = join(__dirname, '/test/game1.txt');
const file = readFileSync(testFilePath, { encoding: 'utf8' });
const input = file.split(/\r?\n/);

const heroInput = input[0].split(",");
const heroName = heroInput[0];
const heroHealth = parseInt(heroInput[1], 10);
const heroAttack = parseInt(heroInput[2], 10);
const heroSpeed = parseInt(heroInput[3], 10);
const hero = new Hero(heroAttack, heroHealth, heroSpeed, heroName);

const maze = new Maze(hero);

const roomsInput = input.slice(1, input.length);
roomsInput.forEach((row) => {
  const elements = row.split(',');
  const roomName = elements[0];
  const monsterName = elements[1];
  const monsterHealth = parseInt(elements[2], 10);
  const monsterAttack = parseInt(elements[3], 10);
  const monsterSpeed = parseInt(elements[4], 10);
  const monsterSpeedDamage = parseInt(elements[5], 10);
  const monsterCloneable = elements[6] === 'true' ? true : false;
  const itemName = elements[7];
  const itemHealth = parseInt(elements[8], 10);
  const itemAttack = parseInt(elements[9], 10);
  const itemSpeed = parseInt(elements[10], 10);

  const item = maze.createItem(itemName, itemAttack, itemHealth, itemSpeed);
  const monster = maze.createMonster(monsterAttack, monsterHealth, monsterSpeed, monsterSpeedDamage, monsterName, monsterCloneable);
  maze.createRoom(roomName, monster, item);
});

maze.startFight();




