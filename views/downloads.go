package views

import (
	"fmt"
	"media-download-manager/types"
	"media-download-manager/views/components"
	"strings"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Downloads(downloads []types.Download) g.Node {
	return Div(
		ID("downloads"),
		g.Attr("hx-get", "/downloads"),
		g.Attr("hx-trigger", "every 1s"),
		g.Attr("hx-swap", "outerHTML"),
		noDownloadsBox(downloads),
		downloadsList(downloads),
	)
}

func noDownloadsBox(downloads []types.Download) g.Node {
	tailwindClasses := []string{
		"list-toggle",
		"p-2",
		"w-full",
		"divide-y",
		"divide-gray-200",
		"border-2",
		"border-solid",
		"rounded-lg",
		"border-gray-200",
		"text-center",
		"font-bold",
		"text-gray-800",
		"dark:text-white",
	}

	if len(downloads) > 0 {
		tailwindClasses = append(tailwindClasses, "hidden")
	}

	return Div(
		Class(strings.Join(tailwindClasses, " ")),
		g.Attr("_", "on checkList set lis to <li/> then if lis.length is 0 remove .hidden from me else add .hidden to me"),
		g.Text("No Downloads"),
	)
}

func downloadsList(downloads []types.Download) g.Node {
	tailwindClasses := []string{
		"list-toggle",
	}

	if len(downloads) == 0 {
		tailwindClasses = append(tailwindClasses, "hidden")
	}

	return Div(
		Class(strings.Join(tailwindClasses, " ")),
		g.Attr("_", "on checkList set lis to <li/> then if lis.length is 0 add .hidden to me else remove .hidden from me"),
		Ul(
			ID("download-list"),
			g.Attr("hx-confirm", "Are you sure?"),
			g.Attr("hx-target", "closest li"),
			g.Attr("hx-swap", "outerHTML"),
			Class("list-toggle w-full divide-y divide-gray-200 border-2 border-solid rounded-lg border-gray-200"),
			g.Attr("_", "on htmx:afterSwap trigger checkList on .list-toggle"),
			g.Group(g.Map(downloads, DownloadElement)),
		),
	)
}

func DownloadElement(download types.Download) g.Node {
	return Li(
		Class("p-3 flex justify-between items-center gap-2"),
		Div(
			Class("flex gap-2 items-center"),
			components.ProgressCircle(download),
			g.If(
				download.Status != types.PENDING,
				Div(
					Div(
						Class("font-bold text-gray-800 dark:text-white"),
						g.Text(download.Title),
					),
					Div(
						Class("text-gray-800 dark:text-white"),
						g.Text(download.TimeRemaining.String()),
					),
				),
			),
		),
		Div(
			Class("text-black dark:text-white fill-current w-6 h-6 cursor-pointer shrink-0"),
			g.Attr("hx-delete", fmt.Sprintf("downloads/%d", download.Id)),
			components.CloseIcon(),
		),
	)
}
