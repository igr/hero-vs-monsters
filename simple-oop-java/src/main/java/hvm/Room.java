package hvm;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

/**
 * Simple record class to hold room data, with custom ctor.
 */
public record Room(String name, List<Monster> monsters, Item item) {
	public Room(String name, Monster monster, Item item) {
		// Sometimes, Java really lack of syntactic sugar.
		this(name, new ArrayList<>(Collections.singletonList(monster)), item);
	}
}
