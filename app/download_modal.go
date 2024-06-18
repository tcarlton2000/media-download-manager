package app

import (
	"log"
	"media-download-manager/views"
	"net/http"
	"os"
	"strings"
)

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

	_ = views.DownloadModal(currentDirectory, getPreviousDirectory(currentDirectory), directories).Render(w)
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

	_ = views.DirectoryPicker(directory, getPreviousDirectory(directory), directories).Render(w)
}

func getPreviousDirectory(directory string) string {
	directories := strings.Split(directory, "/")
	lastDirRemoved := directories[:len(directories)-1]
	return strings.Join(lastDirRemoved, "/")
}
