package app

import (
	"errors"
	"net/http"
	"strconv"
	"text/template"

	"media-download-manager/modules"
)

var DOWNLOADS_TEMPLATE []string = []string{
	"templates/index.html",
	"templates/progress.html",
	"templates/close.html",
}

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

func createDownloadRow(d modules.Download) DownloadRow {
	return DownloadRow{Download: d, ProgressProps: ProgressProps{Progress: d.Progress}}
}

func (a *App) DownloadList(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(DOWNLOADS_TEMPLATE...))

	var downloadRows []DownloadRow
	for _, d := range a.mock.Downloads {
		downloadRows = append(downloadRows, createDownloadRow(d))
	}

	keyMap := map[string][]DownloadRow{
		"Downloads": downloadRows,
	}
	tmpl.Execute(w, keyMap)
}

func (a *App) NewDownload(w http.ResponseWriter, r *http.Request) {
	url := r.PostFormValue("url")
	downloadPath := r.PostFormValue("directory")
	newDownload := modules.DownloadVideo(a.mock, url, downloadPath)
	tmpl := template.Must(template.ParseFiles(DOWNLOADS_TEMPLATE...))
	tmpl.ExecuteTemplate(w, "download-list-element", createDownloadRow(newDownload))
}

func (a *App) DeleteDownload(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	index, err := a.findDownload(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	a.mock.Downloads = append(a.mock.Downloads[:index], a.mock.Downloads[index+1:]...)
}

func (a *App) findDownload(id int) (int, error) {
	for i, d := range a.mock.Downloads {
		if d.Id == id {
			return i, nil
		}
	}

	return 0, errors.New("No download found with id")
}
