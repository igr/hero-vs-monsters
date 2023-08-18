package hvm;

import java.io.IOException;

public class HeroAndMonsters {

	public static void main(String[] args) throws IOException {
		new GameLoader().loadGame().play();
	}
}
