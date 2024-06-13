package modules

import (
	"fmt"
	"media-download-manager/db"
	"media-download-manager/types"
	"os/exec"
)

// Starts download and creates download entry in the database.
func DownloadVideo(db *db.Database, url string, downloadPath string) (types.Download, error) {
	var err error
	download := types.Download{Url: url, DownloadPath: downloadPath}
	download.Id, err = db.NewDownload(download)
	if err != nil {
		return types.Download{}, err
	}

	go startDownload(db, download)
	return download, nil
}

func startDownload(db *db.Database, download types.Download) {
	cmd := exec.Command("yt-dlp", "--no-mtime", download.Url)
	cmd.Dir = download.DownloadPath
	cmd.Stdout = YoutubeDlParser{db: db, download: &download}

	if err := cmd.Run(); err != nil {
		fmt.Println("could not run command: ", err)
		download.Status = types.ERROR
		db.UpdateDownload(download)
	}

	download.Progress = 100
	download.Status = types.COMPLETED
	db.UpdateDownload(download)
}
