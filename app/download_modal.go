package app

import (
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

var DOWNLOAD_MODAL_TEMPLATE []string = []string{
	"templates/modal.html",
	"templates/directory_picker.html",
}

type DownloadModalProps struct {
	CurrentDirectory  string
	PreviousDirectory string
	Directories       []os.DirEntry
}

func (a *App) DownloadModal(w http.ResponseWriter, r *http.Request) {
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	files, err := os.ReadDir(currentDirectory)
	if err != nil {
		log.Fatal(err)
	}

	var directories []os.DirEntry
	for _, file := range files {
		if file.IsDir() {
			directories = append(directories, file)
		}
	}

	tmpl := template.Must(template.ParseFiles(DOWNLOAD_MODAL_TEMPLATE...))
	tmpl.Execute(w, DownloadModalProps{
		CurrentDirectory:  currentDirectory,
		PreviousDirectory: getPreviousDirectory(currentDirectory),
		Directories:       directories,
	})
}

func (a *App) RefreshDirectoryList(w http.ResponseWriter, r *http.Request) {
	directory := r.PostFormValue("directory-picker")
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Print(err)
		return
	}

	var directories []os.DirEntry
	for _, file := range files {
		if file.IsDir() {
			directories = append(directories, file)
		}
	}

	tmpl := template.Must(template.ParseFiles(DOWNLOAD_MODAL_TEMPLATE...))
	tmpl.ExecuteTemplate(w, "directory-list", DownloadModalProps{
		CurrentDirectory:  directory,
		PreviousDirectory: getPreviousDirectory(directory),
		Directories:       directories,
	})
}

func getPreviousDirectory(directory string) string {
	directories := strings.Split(directory, "/")
	lastDirRemoved := directories[:len(directories)-1]
	return strings.Join(lastDirRemoved, "/")
}
