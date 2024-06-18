package components

import (
	"fmt"
	"media-download-manager/types"
	"strings"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func ProgressCircle(download types.Download) g.Node {
	svgClasses := []string{"size-full"}
	innerCircleClasses := []string{"stroke-current"}

	switch download.Status {
	case types.PENDING:
		svgClasses = append(svgClasses, "animate-spin")
		innerCircleClasses = append(innerCircleClasses, "text-blue-600", "dark:text-blue-500")
	case types.COMPLETED:
		innerCircleClasses = append(innerCircleClasses, "text-green-600", "dark:text-green-500")
	case types.ERROR:
		innerCircleClasses = append(innerCircleClasses, "text-red-600", "dark:text-red-500")
	default:
		innerCircleClasses = append(innerCircleClasses, "text-blue-600", "dark:text-blue-500")
	}

	return Div(
		Class("relative size-14 shrink-0"),
		SVG(
			Class(strings.Join(svgClasses, " ")),
			g.If(
				download.Status != types.PENDING,
				Style("transition: width .6 ease"),
			),
			Width("18"),
			Height("18"),
			g.Attr("viewBox", "0 0 18 18"),
			g.Attr("xmlns", "http://www.w3.org/2000/svg"),
	
			// Background Circle
			circle(
				Class("stroke-current text-gray-200 dark:text-gray-700"),
				g.Attr("stroke-width", "2"),
			),
	
			// Progress Circle inside a group with rotation
			g.El(
				"g",
				Class("origin-center -rotate-90 transform"),
				circle(
					Class(strings.Join(innerCircleClasses, " ")),
					g.Attr("stroke-width", "2"),
					g.Attr("stroke-dasharray", "43.96"),
					dashOffset(download),
				),
			),
		),

		// Percentage Text
		g.If(
			download.Status != types.PENDING,
			Div(
				Class("absolute top-1/2 start-1/2 transform -translate-y-1/2 -translate-x-1/2 items-center"),
				Span(
					Class("text-center text-xs font-bold text-gray-800 dark:text-white"),
					g.Text(fmt.Sprintf("%.0f%%", download.Progress)),
				),
			),
		),
	)
}

func circle(children ...g.Node) g.Node {
	return g.El(
		"circle",
		append(
			children,
			g.Attr("cx", "9"),
			g.Attr("cy", "9"),
			g.Attr("r", "7"),
			g.Attr("fill", "none"),
		)...,
	)
}

func dashOffset(download types.Download) g.Node {
	var dashOffsetValue float32
	if download.Status == types.PENDING {
		dashOffsetValue = 35
	} else {
		dashOffsetValue = ((100 - download.Progress) / 100) * 43.96
	}

	return g.Attr("stroke-dashoffset", fmt.Sprintf("%.2f", dashOffsetValue))
}
