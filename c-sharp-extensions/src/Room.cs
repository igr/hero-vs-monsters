namespace HeroVsMonsters;

public record Room(string Name, List<Monster> Monsters, Item Item)
{
    public List<Monster> AliveMonsters() => Monsters.Where(m => m.IsAlive).ToList();
}