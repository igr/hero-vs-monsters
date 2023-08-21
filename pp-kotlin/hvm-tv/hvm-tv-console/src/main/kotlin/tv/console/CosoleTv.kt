package tv.console

import tv.ShowOnTv

object ConsoleTv: ShowOnTv {
	override fun invoke(message: String) = println(message)
}