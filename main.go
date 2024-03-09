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

	http.HandleFunc("/", app.Index)
	http.HandleFunc("GET /downloads", app.DownloadList)
	http.HandleFunc("GET /modal", app.DownloadModal)
	http.HandleFunc("POST /new-download/", app.NewDownload)
	http.HandleFunc("DELETE /downloads/{id}", app.DeleteDownload)
	http.HandleFunc("POST /directories/", app.RefreshDirectoryList)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
