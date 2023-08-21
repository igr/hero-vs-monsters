package hvm.app

import tv.console.ConsoleTv

fun main() {
    val game = loadGame()
    val events = play(game)
    televiseGameEvents(events, ConsoleTv)
}