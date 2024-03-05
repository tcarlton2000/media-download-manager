package modules

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/kkdai/youtube/v2"
)

type Download struct {
	Id            int
	Title         string
	DownloadPath  string
	Url           string
	Status        Status
	TimeRemaining time.Duration
	Progress      float32
}

func DownloadVideo(m Mock, url string, downloadPath string) Download {
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

	return m.NewDownload(video.Title, url, downloadPath)
}
