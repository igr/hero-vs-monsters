package hvm;

import java.io.IOException;

public class HeroAndMonsters {

	public static void main(String[] args) throws IOException {
		var game = new GameLoader().loadGame();
		game.play();
	}
}
