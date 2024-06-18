package components

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func CloseIcon() g.Node {
	return SVG(
		g.Attr("xmlns", "http://www.w3.org/2000/svg"),
		g.Attr("x", "0px"),
		g.Attr("y", "0px"),
		g.Attr("viewBox", "0 0 50 50"),
		g.El(
			"path",
			g.Attr("d", "M 7.71875 6.28125 L 6.28125 7.71875 L 23.5625 25 L 6.28125 42.28125 L 7.71875 43.71875 L 25 26.4375 L 42.28125 43.71875 L 43.71875 42.28125 L 26.4375 25 L 43.71875 7.71875 L 42.28125 6.28125 L 25 23.5625 Z"),
		),
	)
}
