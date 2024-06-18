package views

import (
	"media-download-manager/types"
	"media-download-manager/views/components"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func IndexView(downloads []types.Download) g.Node {
	return Doctype(
		HTML(
			Lang("en"),
			Head(
				Meta(Charset("UTF-8")),
				Meta(
					Name("viewport"),
					Content("width=device-width, initial-scale=1.0"),
				),
				Link(Href("/static/css/modal.css"), Rel("stylesheet")),
				Link(Href("/static/css/output.css"), Rel("stylesheet")),
				Script(Src("https://unpkg.com/htmx.org@1.9.11")),
				Script(Src("https://unpkg.com/hyperscript.org@0.9.12")),
				TitleEl(g.Text("Media Download Manager")),
			),
			Body(
				Class("bg-white dark:bg-gray-700"),
				Div(
					Class("container mx-auto px-4"),
					pageHeader(),
					newDownloadButton(),
					Downloads(downloads),
				),
			),
		),
	)
}

func pageHeader() g.Node {
	return H1(
		Class("mb-4 text-4xl font-bold text-gray-800 dark:text-white"),
		g.Text("Media Download Manager"),
	)
}

func newDownloadButton() g.Node {
	return Div(
		Class("w-full text-right mb-4"),
		components.PrimaryButton(
			g.Attr("hx-get", "/modal"),
			g.Attr("hx-target", "body"),
			g.Attr("hx-swap", "beforeend"),
			g.Text("New Download"),
		),
	)
}
