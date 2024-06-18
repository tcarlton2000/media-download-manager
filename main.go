package main

import (
	"fmt"
	"log"
	"net/http"

	"media-download-manager/app"
)

func main() {
	app := app.App{}
	app.Init()
	fmt.Println("App Started")

	http.HandleFunc("/", app.BasicAuth(app.Index))
	http.HandleFunc("GET /downloads", app.BasicAuth(app.DownloadList))
	http.HandleFunc("GET /modal", app.BasicAuth(app.DownloadModal))
	http.HandleFunc("POST /new-download", app.BasicAuth(app.NewDownload))
	http.HandleFunc("DELETE /downloads/{id}", app.BasicAuth(app.DeleteDownload))
	http.HandleFunc("POST /directories", app.BasicAuth(app.RefreshDirectoryList))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
