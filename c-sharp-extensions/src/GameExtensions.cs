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

    public static IEnumerable<(Room room, List<Monster> survivedMonsters)> GetSurvivedMonstersInRoomOrUseItem(
        this IEnumerable<Room> rooms, Hero hero, ConsoleTv tv)
    {
        foreach (var room in rooms.Where(_ => hero.IsAlive()))
        {
            while (room.Monsters.Any(m => m.IsAlive()))
            {
                yield return (room, room.Monsters.Where(m => m.IsAlive()).ToList());
            }
            
            if (hero.IsAlive())
            {
                hero.UseItem(room.Item);    
                tv.Show("Hero " + hero.Name + " founds " + room.Item.Name);
            }
        }
    }

    public static IEnumerable<(Room room, List<Monster> survivedMonsters)> CloneSurvivedMonstersInRoom(
        this IEnumerable<(Room room, List<Monster> survivedMonsters)> survivedMonstersInRoom, ConsoleTv tv)
    {
        foreach (var (room, survivedMonsters) in survivedMonstersInRoom)
        {
                var clonedMonsters = new List<Monster>();
                survivedMonsters.ForEach(m => clonedMonsters.AddRange(m.TrySpawnClone()));
            
                if (clonedMonsters.Any())
                {
                    tv.Show($"Monster {survivedMonsters.First().Name} cloned!");
                    room.Monsters.AddRange(clonedMonsters);
                }
                
                yield return (room, room.Monsters.Where(m => m.IsAlive()).ToList());
        }
    }

    public static void FightWithSurvivedMonstersInRoom(
        this IEnumerable<(Room room, List<Monster> survivedMonsters)> survivedMonstersInRoom,
        Action<Monster> fightMonster)
    {
        foreach (var (_, survivedMonsters) in survivedMonstersInRoom)
        {
            survivedMonsters.ForEach(fightMonster);
        }
    }
}