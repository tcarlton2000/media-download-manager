package app

import (
	"net/http"
	"strconv"

	"media-download-manager/modules"
	"media-download-manager/views"
)

func (a *App) Index(w http.ResponseWriter, r *http.Request) {
	downloads, err := a.db.GetDownloads()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = views.IndexView(downloads).Render(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (a *App) DownloadList(w http.ResponseWriter, r *http.Request) {
	downloads, err := a.db.GetDownloads()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = views.Downloads(downloads).Render(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (a *App) NewDownload(w http.ResponseWriter, r *http.Request) {
	url := r.PostFormValue("url")
	downloadPath := r.PostFormValue("directory")
	newDownload, err := modules.DownloadVideo(a.db, url, downloadPath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = views.DownloadElement(newDownload).Render(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
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
