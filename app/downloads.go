package app

import (
	"net/http"
	"strconv"
	"text/template"

	"media-download-manager/modules"
	"media-download-manager/types"
)

var DOWNLOADS_TEMPLATE []string = []string{
	"templates/index.html",
	"templates/progress.html",
	"templates/close.html",
}

type DownloadRow struct {
	Download      types.Download
	ProgressProps ProgressProps
}

type ProgressProps struct {
	Progress float32
	Status   types.Status
}

func (p ProgressProps) DashOffset() float32 {
	// The last number needs to match the "stroke-dasharray" in progress.html.
	return ((100 - p.Progress) / 100) * 43.96
}

func createDownloadRow(d types.Download) DownloadRow {
	return DownloadRow{
		Download: d,
		ProgressProps: ProgressProps{
			Progress: d.Progress,
			Status:   d.Status,
		},
	}
}

func (a *App) Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(DOWNLOADS_TEMPLATE...))
	downloads, err := a.db.GetDownloads()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	var downloadRows []DownloadRow
	for _, d := range downloads {
		downloadRows = append(downloadRows, createDownloadRow(d))
	}

	keyMap := map[string][]DownloadRow{
		"Downloads": downloadRows,
	}
	tmpl.Execute(w, keyMap)
}

func (a *App) DownloadList(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(DOWNLOADS_TEMPLATE...))
	downloads, err := a.db.GetDownloads()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	var downloadRows []DownloadRow
	for _, d := range downloads {
		downloadRows = append(downloadRows, createDownloadRow(d))
	}

	keyMap := map[string][]DownloadRow{
		"Downloads": downloadRows,
	}
	tmpl.ExecuteTemplate(w, "downloads", keyMap)
}

func (a *App) NewDownload(w http.ResponseWriter, r *http.Request) {
	url := r.PostFormValue("url")
	downloadPath := r.PostFormValue("directory")
	newDownload, err := modules.DownloadVideo(a.db, url, downloadPath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles(DOWNLOADS_TEMPLATE...))
	tmpl.ExecuteTemplate(w, "download-list-element", createDownloadRow(newDownload))
}

func (a *App) DeleteDownload(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = a.db.DeleteDownload(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
