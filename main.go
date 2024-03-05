package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/kkdai/youtube/v2"
)

type Status int

const (
	PENDING     = Status(0)
	IN_PROGRESS = Status(1)
	COMPLETED   = Status(2)
	ERROR       = Status(3)
)

type DownloadModalProps struct {
	CurrentDirectory  string
	PreviousDirectory string
	Directories       []os.DirEntry
}

type ProgressProps struct {
	Progress float32
}

func (p ProgressProps) DashOffset() float32 {
	return ((100 - p.Progress) / 100) * 43.96
}

func getPreviousDirectory(directory string) string {
	directories := strings.Split(directory, "/")
	lastDirRemoved := directories[:len(directories)-1]
	return strings.Join(lastDirRemoved, "/")
}

func downloadVideo(id int, url string) Download {
	client := youtube.Client{}

	log.Print("Getting video info...")
	video, err := client.GetVideo(url)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Getting video formats...")
	formats := video.Formats.WithAudioChannels()

	log.Print("Getting video stream...")
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	log.Print("Creating video file...")
	file, err := os.Create("video.mp4")
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	// This part goes into the thread
	log.Print("Copying video to file...")
	_, err = io.Copy(file, stream)
	if err != nil {
		log.Fatal(err)
	}

	return Download{Id: id, Title: video.Title, Url: url, Status: PENDING, Progress: 0, TimeRemaining: time.Duration(0)}
}

func main() {
	timeStrings := []string{"00m46s", "01m12s", "06m59s", "12m45s"}

	var timeDurations []time.Duration
	for _, ts := range timeStrings {
		d, err := time.ParseDuration(ts)

		if err != nil {
			log.Fatalf("Time Duration parse failed: %v", err)
		}

		timeDurations = append(timeDurations, d)
	}

	downloads := []Download{
		{Id: 0, Title: "Download 0", Url: "https://youtube.com/0", Status: COMPLETED, Progress: 100, TimeRemaining: timeDurations[0]},
		{Id: 1, Title: "Download 1", Url: "https://youtube.com/1", Status: ERROR, Progress: 67, TimeRemaining: timeDurations[1]},
		{Id: 2, Title: "Download 2", Url: "https://youtube.com/2", Status: IN_PROGRESS, Progress: 50, TimeRemaining: timeDurations[2]},
		{Id: 3, Title: "Download 3", Url: "https://youtube.com/3", Status: PENDING, Progress: 0, TimeRemaining: timeDurations[3]},
	}
	nextId := 4

	fmt.Println("hello world")

	modal := func(w http.ResponseWriter, r *http.Request) {
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

		tmpl := template.Must(template.ParseFiles("templates/modal.html"))
		tmpl.Execute(w, DownloadModalProps{
			CurrentDirectory:  currentDirectory,
			PreviousDirectory: getPreviousDirectory(currentDirectory),
			Directories:       directories,
		})
	}

	dir := func(w http.ResponseWriter, r *http.Request) {
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

		tmpl := template.Must(template.ParseFiles("templates/modal.html"))
		tmpl.ExecuteTemplate(w, "directory-list", DownloadModalProps{
			CurrentDirectory:  directory,
			PreviousDirectory: getPreviousDirectory(directory),
			Directories:       directories,
		})
	}

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html", "templates/progress.html"))

		var downloadRows []DownloadRow
		for _, d := range downloads {
			downloadRows = append(downloadRows, d.DownloadRow())
		}

		keyMap := map[string][]DownloadRow{
			"Downloads": downloadRows,
		}
		tmpl.Execute(w, keyMap)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		url := r.PostFormValue("url")
		newDownload := downloadVideo(nextId, url)
		downloads = append(downloads, newDownload)
		tmpl := template.Must(template.ParseFiles("templates/index.html", "templates/progress.html"))
		tmpl.ExecuteTemplate(w, "download-list-element", newDownload.DownloadRow())
		nextId++
	}

	delete := func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("DELETE")
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/modal", modal)
	http.HandleFunc("/new-download/", h2)
	http.HandleFunc("/downloads", delete)
	http.HandleFunc("/directories/", dir)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
