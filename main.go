package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tcarlton2000/media-download-manager/app"
)

func main() {
	app := app.App{}
	app.Init()
	fmt.Println("App Started")

	delete := func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("DELETE")
	}

	http.HandleFunc("/", app.DownloadList)
	http.HandleFunc("/modal", app.DownloadModal)
	http.HandleFunc("/new-download/", app.NewDownload)
	http.HandleFunc("/downloads", delete)
	http.HandleFunc("/directories/", app.RefreshDirectoryList)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
