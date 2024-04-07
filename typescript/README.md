# Hero vs Monsters (Node.js + Typescript)

## Prerequisites 
To start app you need to have node.js installed on your system trough terminal.

Suggested version of node.js is v20, but it should work ok with v16+.

Let's test if node works on your terminal with command:

`node -v`

if your output is something similar like below, we are good to continue.

> v20.12.1 (v16.x.x/v18.x.x/v20.x.x)

## Start node.js app

To start app follow next 2 commands, that will install all dependencies and start game

`npm install`

`npm start`

Like this, game will start with default input from file located in `input/game1.txt` 

Change this file with rules below to change game

```
<Hero>,<health>,<attack>,<speed>
<Room A>,<Monster A>,<health>,<attack>,<speed>,<speed damage>,<clonable>,<Item A>,<health>,<attack>,<speed>
<Room B>,<Monster B>,<health>,<attack>,<speed>,<speed damage>,<clonable>,<Item B>,<health>,<attack>,<speed>
```

# Details about game and task requests are below!

## 🤴⚔️😈 The problem

Implement gameplay for a simple RPG. In the game, there is the **hero**, our main character, and the **maze**, which acts as the game map.

The game map (maze) is a list of **rooms**. The hero moves from one room to another, never returning to a room he has already visited. There is always one exit from the room, the hero automatically proceeds through it.

In each room, there is a **monster** and an **item**. The hero begins his journey from the first room, battling a monster in each room along his path. Once a monster is defeated, the hero picks up an item in the same room to enhance himself.

Both heroes and monsters posses three attributes:

- attack damage,
- health, and
- speed.

Items have the same attributes, used to enhance the hero, after the monster is defeated.

Monsters have an another attribute:

- speed damage

The hero and monster take turns attacking each other. The `attack damage` indicates the amount of damage the hero or the monster will cause to their opponent (i.e., reducing their health) each turn. One is defeated once their health drops to `0` (or less). `Speed` determines which character attacks first (the one with a higher speed). `Speed damage` is the reduction in speed the hero will suffer if hit by the monster. The game is won when the last monster is defeated. Whenever the hero is defeated, it’s game over.

Some monsters have the ability to clone themselves. This happens on the monster's turn when his health is less than `25%` of its initial health. The health is then split between the two. A monster can only clone itself once.

Monster roars before the attack. Roar is constructed by concatenating random elements from the following set:

- single `H`, `W`, `L` letter.
- multiple (3-6, random) `R`, `O`, `A` letters.

The outcome of the game is deterministic, i.e., the same input always produces the same output. There is no randomization in the game flow.

### Non-functional requirements

The maze with monsters is loaded from the input file `game.txt` that looks like this:

```
<Hero>,<health>,<attack>,<speed>
<Room A>,<Monster A>,<health>,<attack>,<speed>,<speed damage>,<clonable>,<Item A>,<health>,<attack>,<speed>
<Room B>,<Monster B>,<health>,<attack>,<speed>,<speed damage>,<clonable>,<Item B>,<health>,<attack>,<speed>
```

The whole battle should be televised :) Every move must be printed to the console. In the future, the console may be replaced with a different output system. What has to be printed:

- the moment hero enters the room,
- the monster's roar,
- the attack sequence between the hero and the monster,
- monster's cloning,
- the outcome of the battle,
- the moment hero picks up the item (if any)
- the final outcome of the game.

Try to use name `Television` for the output system. Example output for the [game1.txt](game1.txt):

```plaintext
Hero Beorn enters Hallway
Monster Haunthand attacks: WHAAAAAARRROOOL
Hero Beorn fights Haunthand
Monster Haunthand attacks: AAAWHLOOOOORRRR
Hero Beorn fights Haunthand
Monster Haunthand is dead
Hero Beorn founds Sword
Hero Beorn enters Dark Room
Monster Helltree attacks: RRRRRHLWOOOOOAAAAA
Hero Beorn fights Helltree
Monster Helltree attacks: AAAHWOOORRRRRL
Hero Beorn fights Helltree
Monster Helltree attacks: WRRROOOOOOHLAAAA
Hero Beorn fights Helltree
Monster Helltree is dead
Hero Beorn founds Shield
Hero Beorn enters Tower
Monster Red Dragon attacks: AAAAAHRRRRRLOOOOOOW
Hero Beorn fights Red Dragon
Monster Red Dragon attacks: WAAAAAHOOORRRRL
Hero Beorn fights Red Dragon
Monster Red Dragon attacks: RRRRROOOOOWLAAAAH
Hero Beorn fights Red Dragon
Monster Red Dragon cloned!
Monster Red Dragon attacks: RRRRRRLOOOOHWAAAAAA
Hero Beorn fights Red Dragon
Monster Red Dragon is dead
Monster Red Dragon clone attacks: WAAAAAHLRRRROOOO
Hero Beorn fights Red Dragon clone
Monster Red Dragon clone is dead
Hero Beorn founds Gold
Hero Beorn wins!
```
