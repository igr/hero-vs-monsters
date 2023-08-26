namespace HeroVsMonsters;

public record struct Hero(string Name, int Health, int Attack, int Speed)
{
    public bool IsAlive => Health > 0;

    public bool IsDead => Health <= 0;

    private void Hit(Monster target) => target.TakeHit(Attack);

    public void TakeHit(int attack, int speedDamage)
    {
        Health -= attack;
        Speed -= speedDamage;
    }

    public void UseItem(Item item)
    {
        Health += item.Health;
        Attack += item.Attack;
        Speed += item.Speed;
    }

    public void Fight(Monster monster)
    {
        if (monster.Speed > Speed)
        {
            monster.Hit(this);
            if (IsAlive)
            {
                Hit(monster);
            }
        }
        else
        {
            Hit(monster);
            if (monster.IsAlive)
            {
                monster.Hit(this);
            }
        }
    }
}