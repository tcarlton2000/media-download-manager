package components

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func PrimaryButton(children ...g.Node) g.Node {
	return Button(
		append(
			children,
			Class("bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"),
		)...,
	)
}

func SecondaryButton(children ...g.Node) g.Node {
	return Button(
		append(
			children,
			Class("font-bold py-2 px-4 rounded text-gray-800 dark:text-white"),
		)...,
	)
}