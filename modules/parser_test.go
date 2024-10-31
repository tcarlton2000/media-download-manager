package modules

import (
	"media-download-manager/db"
	"media-download-manager/types"
	"os"
	"testing"
	"time"
)

func TestParser(t *testing.T) {
	ydp := setupParser(t)

	ydp.Write([]byte("[download] Destination: youtube_file.mp4"))
	expectedDownload := types.Download{
		Id:           ydp.download.Id,
		Status:       types.IN_PROGRESS,
		Title:        "youtube_file.mp4",
		Url:          "youtube.com/1234",
		DownloadPath: "/downloads",
	}
	validateUpdate(t, ydp, expectedDownload)

	ydp.Write([]byte("[download]   0.0% of   54.03MiB at  Unknown B/s ETA Unknown"))
	validateUpdate(t, ydp, expectedDownload)

	ydp.Write([]byte("[download]   0.0% of   54.03MiB at    2.36MiB/s ETA 00:23"))
	expectedDownload.TimeRemaining, _ = time.ParseDuration("00m23s")
	validateUpdate(t, ydp, expectedDownload)

	ydp.Write([]byte("[download]   35.2% of   54.03MiB at    2.36MiB/s ETA 00:09"))
	expectedDownload.Progress = 35.2
	expectedDownload.TimeRemaining, _ = time.ParseDuration("00m09s")
	validateUpdate(t, ydp, expectedDownload)

	ydp.Write([]byte("[download]   100% of   54.03MiB at    2.36MiB/s ETA 00:00"))
	expectedDownload.TimeRemaining, _ = time.ParseDuration("00m00s")
	expectedDownload.Progress = 100.0
	validateUpdate(t, ydp, expectedDownload)

	teardown()
}

func TestParser_fragmentedFile(t *testing.T) {
	ydp := setupParser(t)
	ydp.Write([]byte("[download]  26.3% of ~ 129.36MiB at    2.18MiB/s ETA 00:48 (frag 11/42)"))

	expectedTimeRemaining, _ := time.ParseDuration("00m48s")
	expectedDownload := types.Download{
		Id:            ydp.download.Id,
		Status:        types.PENDING,
		Title:         "",
		Url:           "youtube.com/1234",
		DownloadPath:  "/downloads",
		Progress:      26.3,
		TimeRemaining: expectedTimeRemaining,
	}
	validateUpdate(t, ydp, expectedDownload)

	teardown()
}

func setupParser(t *testing.T) *YoutubeDlParser {
	d := types.Download{
		Status:       types.PENDING,
		Url:          "youtube.com/1234",
		DownloadPath: "/downloads",
	}
	os.Remove("test/media-download-manager.db")
	os.Mkdir("test", 0777)
	db := db.OpenDb("test")

	var err error
	d.Id, err = db.NewDownload(d)

	if err != nil {
		t.Errorf("Error: %q", err)
	}

	return &YoutubeDlParser{download: &d, db: db}
}

func teardown() {
	os.Remove("test/media-download-manager.db")
	os.Remove("test")
}

func validateUpdate(t *testing.T, ydp *YoutubeDlParser, expectedDownload types.Download) {
	downloads, err := ydp.db.GetDownloads()
	if err != nil {
		t.Errorf("Error: %q", err)
	}

	if len(downloads) != 1 {
		t.Errorf("Expected 1 download, recevied %d", len(downloads))
	}

	if downloads[0] != expectedDownload {
		t.Errorf("Expected %#v, received %#v", expectedDownload, downloads[0])
	}
}
