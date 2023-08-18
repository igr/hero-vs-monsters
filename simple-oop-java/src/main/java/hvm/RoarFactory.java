package hvm;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Random;
import java.util.stream.Collectors;

/**
 * Encapsulates the logic of generating Monster roars.
 */
public class RoarFactory {

	private static final List<String> roar = "HWLROA".chars().mapToObj(Character::toString).toList();

	public String roar() {
		var rnd = new Random();

		var list = new ArrayList<>(roar);
		Collections.shuffle(list);

		return list.stream().map(c -> {
			switch (c) {
				case "H", "W", "L" -> {
					return c;
				}
				case "R", "O", "A" -> {
					return String.join("", Collections.nCopies(rnd.nextInt(3, 7), c));
				}
				default -> {
					return "";
				}
			}
		}).collect(Collectors.joining());
	}
}
