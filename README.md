# Heros vs Monsters

This repository contains the list of different implementations of the same problem, described below. The purpose is to present different programming styles, languages, methodologies, patterns... and to compare them.

Each subfolder must be a complete solution, with its own build system. The only common part is the problem description.

Please add as many comments as possible to your code, so that it is easy to understand. The purpose is to learn from each other, and to compare different approaches.

## Implementations

TBA

## ⚔️ The problem

Implement gameplay for a simple RPG. In the game, there is the **hero**, our main character, and the **maze**, which acts as the game map.

The game map (maze) is a list of **rooms**. The hero moves from one room to another, never returning to a room he has already visited. There is alway one exit from the room, the hero automatically proceeds through it.

In each room, there is a **monster** and an **item**. The hero begins his journey from the first room, battling a monster in each room along his path. Once a monster is defeated, the hero picks up an item in the same room to enhance himself.

Both heroes and monsters posses three attributes:
+ attack damage,
+ health, and
+ speed.

Monsters have an another attribute:
+ speed damage

The hero and monster take turns attacking each other. The `attack damage` indicates the amount of damage the hero or the monster will cause to their opponent (i.e., reducing their health) each turn. One is defeated once their health drops to `0` (or less). `Speed` determines which character attacks first (the one with a higher speed). `Speed damage` is the amount of damage the hero will suffer if hit by the monster. The game is won when the last monster is defeated. Whenever the hero is defeated, it’s game over.

Some monsters have the ability to clone themselves. This happens on the monster's turn when his health is less than `25%` of its initial health. The health is then split between the two, but the speed remains the same (effectively, it is the clone's turn after the split.) A monster can only clone itself once.

Monster may roar before the attack (2/3 of the time). Roar is constructed by concatenating random elements from the following set: 
+ single `H`, `W`, `L` letter.
+ multiple (3-6) `R`, `O`, `A` letters.

### Non-functional requirements

The maze with monsters is loaded from the input file `game.txt` that looks like this:

```
<Hero>,<attack>,<health>,<speed>
<Room A>,<Monster A>,<attack>,<health>,<speed>,<speed damage>,<clonable>
<Room B>,<Monster B>,<attack>,<health>,<speed>,<speed damage>,<clonable>
```

The whole battle should be televised :) Every move must be printed to the console. In the future, the console may be replaced with a different output system.
