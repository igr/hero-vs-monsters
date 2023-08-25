namespace HeroVsMonsters;

public class GameLoader
{
    public Game LoadGame()
    {
        var gameLines = File.ReadAllLines(Path.Combine(Directory.GetParent(Environment.CurrentDirectory)!.FullName, "game1.txt"));
        var hero = ParseHero(gameLines[0]);
        var rooms = gameLines.Skip(1).Select(ParseRoom).ToList();
        return new Game(hero, rooms, new ConsoleTv());
    }

    private Hero ParseHero(string line)
    {
        var heroParts = line.Split(",");
        var name = heroParts[0].Trim();
        var health = int.Parse(heroParts[1]);
        var attack = int.Parse(heroParts[2]);
        var speed = int.Parse(heroParts[3]);
        return new Hero(name, health, attack, speed);
    }

    private Room ParseRoom(string line)
    {
        var roomParts = line.Split(",");
        var roomName = roomParts[0].Trim();
        var monster = ParseMonster(roomParts);
        var item = ParseItem(roomParts);
        return new Room(roomName, new List<Monster> { monster }, item);
    }

    private Monster ParseMonster(string[] lines)
    {
        var name = lines[1].Trim();
        var health = int.Parse(lines[2]);
        var attack = int.Parse(lines[3]);
        var speed = int.Parse(lines[4]);
        var speedDamage = int.Parse(lines[5]);
        var duplicable = bool.Parse(lines[6]);
        return new Monster(name, health, attack, speed, speedDamage, duplicable);
    }

    private Item ParseItem(string[] lines)
    {
        var name = lines[7].Trim();
        var health = int.Parse(lines[8]);
        var attack = int.Parse(lines[9]);
        var speed = int.Parse(lines[10]);
        return new Item(name, health, attack, speed);
    }
}