package modules

import (
	"context"
	"fmt"
	"io"
	"log"
	"media-download-manager/db"
	"media-download-manager/types"
	"os"

	"github.com/wader/goutubedl"
)

func DownloadVideo(db *db.Database, url string, downloadPath string) (types.Download, error) {
	var err error
	download := types.Download{Url: url, DownloadPath: downloadPath}
	download.Id, err = db.NewDownload(download)
	if err != nil {
		return types.Download{}, err
	}

	go finishDownload(db, download)
	return download, nil
}

func finishDownload(db *db.Database, download types.Download) {
	log.Print("Starting Info...")
	result, err := goutubedl.New(context.Background(), download.Url, goutubedl.Options{})
	if err != nil {
		log.Print(err)
		download.Status = types.ERROR
		db.UpdateDownload(download)
		return
	}

	download.Status = types.IN_PROGRESS
	download.Title = result.Info.Title
	log.Print(result.Info.Title)
	db.UpdateDownload(download)

	log.Print("Downloading Video...")
	downloadResult, err := result.Download(context.Background(), "best")
	if err != nil {
		handleDownloadError(db, download, err)
		return
	}
	defer downloadResult.Close()

	log.Print("Creating File...")
	f, err := os.Create(fmt.Sprintf("%s/%s.mp4", download.DownloadPath, download.Title))
	if err != nil {
		handleDownloadError(db, download, err)
		return
	}
	defer f.Close()

	download.Status = types.IN_PROGRESS
	db.UpdateDownload(download)

	log.Print("Copying file...")
	_, err = io.Copy(f, downloadResult)
	if err != nil {
		handleDownloadError(db, download, err)
		return
	}

	download.Progress = 100.0
	download.Status = types.COMPLETED
	db.UpdateDownload(download)
}

func handleDownloadError(db *db.Database, d types.Download, err error) {
	log.Print(err)
	d.Status = types.ERROR
	db.UpdateDownload(d)
}
