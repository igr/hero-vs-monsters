import { join } from 'node:path';
import { readFileSync } from 'node:fs';
import { Item, Hero, Maze, Monster, Room } from './entities';
import { logger } from "./logger";


logger("🚀🚀🚀 Initializing hero-vs-monsters in Node.js + Typescript");

const testFilePath = join(__dirname, '/input/game1.txt');
const file = readFileSync(testFilePath, { encoding: 'utf8' });
const input = file.split(/\r?\n/);

const heroInput = input[0].split(",");

const hero = new Hero({
  name: heroInput[0],
  health: parseInt(heroInput[1], 10),
  attackDamage: parseInt(heroInput[2], 10),
  speed: parseInt(heroInput[3], 10)
});

const maze = new Maze(hero);

const roomsInput = input.slice(1, input.length);

roomsInput.forEach((row) => {
  const elements = row.split(',');
  if (elements.length !== 11) return;

  const monster = new Monster({
    name: elements[1],
    health: parseInt(elements[2], 10),
    attackDamage: parseInt(elements[3], 10),
    speed: parseInt(elements[4], 10),
    speedDamage: parseInt(elements[5], 10),
    cloneable: elements[6] === 'true' ? true : false
  },
  );

  const item = new Item({
    name: elements[7],
    health: parseInt(elements[8], 10),
    attackDamage: parseInt(elements[9], 10),
    speed: parseInt(elements[10], 10)
  });

  const roomName = elements[0];
  const room = new Room(roomName, monster, item);
  maze.addRoom(room);
});

maze.startFight();




