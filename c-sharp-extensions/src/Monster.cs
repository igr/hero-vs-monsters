namespace HeroVsMonsters;

public class Monster
{
    public readonly string Name;
    public readonly int Speed;
    
    private readonly int _initialHealth;
    private readonly int _attack;
    private readonly int _speedDamage;
    
    private bool _cloneable;
    private int _health;
    
    public Monster(string name, int health, int attack, int speed, int speedDamage, bool cloneable)
    {
        Name = name;
        Speed = speed;
        _health = health;
        _initialHealth = health;
        _attack = attack;
        _speedDamage = speedDamage;
        _cloneable = cloneable;
    }
    
    public bool IsAlive => _health > 0;

    public bool IsDead => !IsAlive;
    
    public void Hit(Hero target) => target.TakeHit(_attack, _speedDamage);

    public void TakeHit(int attack)
    {
        _health -= attack;
    }
    
    public IEnumerable<Monster> TrySpawnClone()
    {
        if (!_cloneable || _health > _initialHealth / 4) {
            return new List<Monster>();
        }
        
        _health /= 2;
        _cloneable = false;
        return new []{new Monster(Name + " clone", _health / 2, _attack, Speed, _speedDamage, false)};
    }

    public string Roar()
    {
        var shuffledLetters = "HWLROA".OrderBy(_ => Guid.NewGuid());
        var random = new Random();

        return string.Join("", shuffledLetters.Select(letter => letter switch
        {
            'H' or 'W' or 'L' => letter.ToString(),
            'R' or 'O' or 'A' => string.Concat(Enumerable.Repeat(letter, random.Next(3, 7))),
            _ => ""
        }));
    }
}