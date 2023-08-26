namespace HeroVsMonsters;

public static class GameExtensions
{
    public static IEnumerable<Room> GetNextRoomOrUseItem(this IEnumerable<Room> rooms, Hero hero, ConsoleTv tv)
    {
        foreach (var room in rooms.Where(_ => hero.IsAlive))
        {
            tv.Show("Hero " + hero.Name + " enters " + room.Name);
            while (room.AliveMonsters().Any() && hero.IsAlive)
            {
                yield return room;
            }
            
            if (hero.IsAlive)
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
                room.AliveMonsters().SelectMany(m => m.TrySpawnClone()).ToList().ForEach(cloned =>
                { 
                    tv.Show($"Monster {room.AliveMonsters().First().Name} cloned!");
                    room.Monsters.Add(cloned);
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