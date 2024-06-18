package views

import (
	"fmt"
	"io/fs"
	"media-download-manager/views/components"

	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func DownloadModal(currentDirectory string, previousDirectory string, directories []fs.DirEntry) g.Node {
	return Div(
		ID("modal"),
		g.Attr("_", "on closeModal add .closing then wait for animationend then remove me"),

		Div(
			Class("modal-underlay"),
			g.Attr("_", "on click trigger closeModal"),
		),
		Div(
			Class("modal-content bg-white dark:bg-gray-800"),
			H1(
				Class("mb-4 text-xl font-bold text-gray-800 dark:text-white"),
				g.Text("New Download"),
			),
			Form(
				g.Attr("hx-post", "/new-download"),
				g.Attr("hx-target", "#download-list"),
				g.Attr("hx-swap", "afterbegin"),
				g.Attr("hx-on::after-request", " if(event.detail.successful) this.reset()"),
				g.Attr("_", "on htmx:afterRequest trigger checkList on .list-toggle"),

				Div(
					Class("mb-4"),
					Label(
						For("download-url"),
						Class("block text-sm font-bold mb-2 text-gray-800 dark:text-white"),
						g.Text("URL"),
					),
					Input(
						Type("text"),
						Name("url"),
						ID("download-title"),
						Class("shadow appearance-none border rounded w-full py-2 px-3 leading-tight focus:outline-none focus:shadow-outline bg-white text-gray-800 dark:bg-gray-800 dark:text-white"),
					),
				),
				Div(
					Class("mb-6"),
					Label(
						For("directory"),
						Class("block text-sm font-bold mb-2 text-gray-800 dark:text-white"),
						g.Text("Directory"),
					),
					DirectoryPicker(currentDirectory, previousDirectory, directories),
				),
				components.PrimaryButton(
					Type("submit"),
					g.Attr("_", "on click trigger closeModal"),
					g.Text("Submit"),
				),
				components.SecondaryButton(
					Type("button"),
					g.Attr("_", "on click trigger closeModal"),
					g.Text("Close"),
				),
			),
		),
	)
}

type DirectoryEntry struct {
	Name string
	Path string
}

func DirectoryPicker(currentDirectory string, previousDirectory string, directories []fs.DirEntry) g.Node {
	entries := []DirectoryEntry{
		{
			Name: "..",
			Path: previousDirectory,
		},
	}

	for _, directory := range directories {
		entries = append(
			entries,
			DirectoryEntry{
				Name: directory.Name(),
				Path: fmt.Sprintf("%s/%s", currentDirectory, directory.Name()),
			},
		)
	}

	return Div(
		ID("directory-list"),
		Input(
			Name("directory"),
			Class("hidden"),
			Value(currentDirectory),
		),
		Span(
			Class("text-gray-800 dark:text-white text-sm italic"),
			g.Text(fmt.Sprintf("Current Directory: %s", currentDirectory)),
		),
		Div(
			Class("relative w-full p-2 bg-white text-gray-800 dark:bg-gray-800 dark:text-white overflow-hidden px-1 border rounded border-gray-200"),
			Div(
				ID("loading"),
				Class("htmx-indicator absolute w-full h-full align-middle text-center"),
				g.Text("Loading..."),
			),
			g.Group(g.Map(entries, directoryPickerEntry)),
		),
	)
}

func directoryPickerEntry(entry DirectoryEntry) g.Node {
	return Div(
		Class("cursor-pointer"),
		g.Attr("hx-vals", fmt.Sprintf("{\"directory-picker\": \"%s\"}", entry.Path)),
		g.Attr("hx-trigger", "click"),
		g.Attr("hx-post", "/directories"),
		g.Attr("hx-target", "#directory-list"),
		g.Attr("hx-indicator", "#loading"),
		g.Attr("hx-swap", "outerHTML"),
		g.Text(entry.Name),
	)
}
