namespace HeroVsMonsters;

public class Game
{
    private Hero _hero;
    private readonly List<Room> _rooms;
    private readonly ConsoleTv _tv;

    public Game(Hero hero, List<Room> rooms, ConsoleTv tv)
    {
        _hero = hero;
        _rooms = rooms;
        _tv = tv;
    }

    /// <summary>
    /// Use of c# extensions methods (chaining) and power of IEnumerable/yield return .NET feature
    /// These extension methods are located in GameExtensions.cs
    /// </summary>
    public void Play()
    {
        _rooms
            .GetNextRoomOrUseItem(_hero, _tv)
            .CloneSurvivedMonstersInRoom(_tv)
            .FightWithSurvivedMonstersInRoom(FightMonster);
        
        _tv.ShowConditionally($"Hero {_hero.Name} wins!", _hero.IsAlive);
    }
    
    private void FightMonster(Monster monster)
    {
        _tv.Show($"Monster {monster.Name} attacks: {monster.Roar()}");
        _tv.Show($"Hero {_hero.Name} fights {monster.Name}");
        _hero.Fight(monster);
        _tv.ShowConditionally($"Hero {_hero.Name} is dead", _hero.IsDead);
        _tv.ShowConditionally($"Monster {monster.Name} is dead", monster.IsDead);
    }
}