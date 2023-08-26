namespace HeroVsMonsters;

public static class GameExtensions
{
    public static IEnumerable<Room> GetNextRoom(this IEnumerable<Room> allRooms, Hero hero, ConsoleTv tv)
    {
        foreach (var room in allRooms.Where(_ => hero.IsAlive()))
        {
            tv.Show("Hero " + hero.Name + " enters " + room.Name);
            yield return room;
        }
    }

    public static IEnumerable<Room> GetSurvivedMonstersInRoomOrUseItem(this IEnumerable<Room> rooms, Hero hero, ConsoleTv tv)
    {
        foreach (var room in rooms)
        {
            while (room.AliveMonsters().Any())
            {
                yield return room;
            }

            if (hero.IsAlive())
            {
                hero.UseItem(room.Item);    
                tv.Show("Hero " + hero.Name + " founds " + room.Item.Name);
            }
        }
    }

    public static IEnumerable<Room> CloneSurvivedMonstersInRoom(this IEnumerable<Room> rooms, ConsoleTv tv)
    {
        foreach (var room in rooms)
        {
                var clonedMonsters = new List<Monster>();
                room.AliveMonsters().ForEach(m => clonedMonsters.AddRange(m.TrySpawnClone()));
            
                clonedMonsters.ForEach(_ =>
                { 
                    tv.Show($"Monster {room.AliveMonsters().First().Name} cloned!");
                    room.Monsters.AddRange(clonedMonsters);
                });

                yield return room;
        }
    }

    public static void FightWithSurvivedMonstersInRoom(this IEnumerable<Room> rooms, Action<Monster> fightMonster)
    {
        foreach (var room in rooms)
        {
            room.AliveMonsters().ForEach(fightMonster);
        }
    }
}