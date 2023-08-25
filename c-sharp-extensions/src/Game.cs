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
    /// Use of c# extensions methods (chaining) and power of IEnumerable .NET feature:
    /// IEnumerable is the return type from an iterator. An iterator is a method that uses the yield return keywords. 
    /// yield return is different from a normal return statement because, while it does return a value from the function, 
    /// it doesn't “close the book” on that function. The next time a value is expected, the function will continue executing 
    /// statements after the yield return until it hits another yield return.
    /// These extension methods are located in GameExtensions.cs
    /// </summary>
    public void Play()
    {
        _rooms
            .GetNextRoom(_hero, _tv)
            .GetSurvivedMonstersInRoomOrUseItem(_hero, _tv)
            .CloneSurvivedMonstersInRoom(_tv)
            .FightWithSurvivedMonstersInRoom(FightMonster);
        
        if (_hero.IsAlive())
        {
            _tv.Show("Hero " + _hero.Name + " wins!");
        }
    }
    
    private void FightMonster(Monster monster)
    {
        var roar = monster.Roar();
        _tv.Show($"Monster {monster.Name} attacks: {roar}");
        _tv.Show($"Hero {_hero.Name} fights {monster.Name}");
        _hero.Fight(monster);

        if (_hero.IsDead())
        {
            _tv.Show($"Hero {_hero.Name} is dead");
        }
        else if (monster.IsDead())
        {
            _tv.Show($"Monster {monster.Name} is dead");
        }
    }
}