namespace HeroVsMonsters;

public class ConsoleTv
{
    public void Show(string message) => Console.WriteLine(message);
    public void ShowConditionally(string message, bool condition)
    {
        if (condition)
            Show(message);
    }
}