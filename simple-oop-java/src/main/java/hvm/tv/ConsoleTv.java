package hvm.tv;

public class ConsoleTv implements Television {
	@Override
	public void show(String message) {
		System.out.println(message);
	}
}
