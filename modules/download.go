package modules

import (
	"io"
	"log"
	"media-download-manager/db"
	"media-download-manager/types"
	"os"

	"github.com/kkdai/youtube/v2"
)

func DownloadVideo(db *db.Database, url string, downloadPath string) (types.Download, error) {
	client := youtube.Client{}

	log.Print("Getting video info...")
	video, err := client.GetVideo(url)
	if err != nil {
		log.Print(err)
		return types.Download{}, err
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

	download := types.Download{Url: url, Title: video.Title, DownloadPath: downloadPath}
	download.Id, err = db.NewDownload(download)
	if err != nil {
		return types.Download{}, err
	}

	return download, nil
}
