package app

import (
	"net/http"
	"text/template"

	"media-download-manager/modules"
)

type DownloadRow struct {
	Download      modules.Download
	ProgressProps ProgressProps
}

type ProgressProps struct {
	Progress float32
}

func (p ProgressProps) DashOffset() float32 {
	// The last number needs to match the "stroke-dasharray" in progress.html.
	return ((100 - p.Progress) / 100) * 43.96
}

func (a *App) DownloadList(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html", "templates/progress.html", "templates/close.html"))

	var downloadRows []DownloadRow
	for _, d := range a.mock.Downloads {
		downloadRows = append(downloadRows, createDownloadRow(d))
	}

	keyMap := map[string][]DownloadRow{
		"Downloads": downloadRows,
	}
	tmpl.Execute(w, keyMap)
}

func createDownloadRow(d modules.Download) DownloadRow {
	return DownloadRow{Download: d, ProgressProps: ProgressProps{Progress: d.Progress}}
}
