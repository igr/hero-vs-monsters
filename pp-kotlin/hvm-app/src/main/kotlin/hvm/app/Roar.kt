package hvm.app

import java.util.*

private val roarChars: List<Char> = "HWLROA".toList()

fun makeRoar(): String {
	val list = roarChars.toMutableList()
	list.shuffle()

	return list.map {
		when (it) {
			'H', 'W', 'L' -> it
			'R', 'O', 'A' -> it.toString().repeat(Random().nextInt(3) + 3)
			else -> ""
		}
	}.joinToString("")
}

