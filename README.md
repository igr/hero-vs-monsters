# Hero vs Monsters

This repository contains the list of different solutions of the same problem, described below. Its goal is to showcase a variety of programming styles, languages, methodologies, and patterns, and to compare them.

Each subfolder must be a complete solution, with its own build system. The only common part is the problem description.

While this is a minor problem, feel free to express your ideas and demonstrate your principles. For example, if you have a separate module with only one file, that's perfectly acceptable here since our focus is to provide an illustrative example.

You are welcome to add more implementations, or enhance existing ones. It's perfectly fine to have multiple implementations in the same language, provided they're sufficiently distinct.

We encourage you to add as many comments to your code as possible to make it easier to understand. Our aim is to learn from each other and compare different approaches.

## Implementations

+ [Procedural Programming, Kotlin](pp-kotlin/README.md)
+ [Object Oriented Programming, Java](simple-oop-java/README.md)
+ [C# implementation](c-sharp-extensions/README.md)

## 🤴⚔️😈 The problem

Implement gameplay for a simple RPG. In the game, there is the **hero**, our main character, and the **maze**, which acts as the game map.

The game map (maze) is a list of **rooms**. The hero moves from one room to another, never returning to a room he has already visited. There is alway one exit from the room, the hero automatically proceeds through it.

In each room, there is a **monster** and an **item**. The hero begins his journey from the first room, battling a monster in each room along his path. Once a monster is defeated, the hero picks up an item in the same room to enhance himself.

Both heroes and monsters posses three attributes:

+ attack damage,
+ health, and
+ speed.

Items have the same attributes, used to enhance the hero, after the monster is defeated.

Monsters have an another attribute:

+ speed damage

The hero and monster take turns attacking each other. The `attack damage` indicates the amount of damage the hero or the monster will cause to their opponent (i.e., reducing their health) each turn. One is defeated once their health drops to `0` (or less). `Speed` determines which character attacks first (the one with a higher speed). `Speed damage` is the amount of damage the hero will suffer if hit by the monster. The game is won when the last monster is defeated. Whenever the hero is defeated, it’s game over.

Some monsters have the ability to clone themselves. This happens on the monster's turn when his health is less than `25%` of its initial health. The health is then split between the two. A monster can only clone itself once.

Monster roars before the attack. Roar is constructed by concatenating random elements from the following set: 
+ single `H`, `W`, `L` letter.
+ multiple (3-6, random) `R`, `O`, `A` letters.

The outcome of the game is deterministic, i.e., the same input always produces the same output. There is no randomization in the game flow.

### Non-functional requirements

The maze with monsters is loaded from the input file `game.txt` that looks like this:

```
<Hero>,<health>,<attack>,<speed>
<Room A>,<Monster A>,<health>,<attack>,<speed>,<speed damage>,<clonable>,<Item A>,<health>,<attack>,<speed>
<Room B>,<Monster B>,<health>,<attack>,<speed>,<speed damage>,<clonable>,<Item B>,<health>,<attack>,<speed>
```

The whole battle should be televised :) Every move must be printed to the console. In the future, the console may be replaced with a different output system. What has to be printed:

+ the moment hero enters the room,
+ the monster's roar,
+ the attack sequence between the hero and the monster,
+ monster's cloning,
+ the outcome of the battle,
+ the moment hero picks up the item (if any)
+ the final outcome of the game.

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