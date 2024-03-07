package modules

import (
	"log"
	"time"
)

type Mock struct {
	NextId    int
	Downloads []Download
}

func (m *Mock) Init() {
	timeStrings := []string{"00m46s", "01m12s", "06m59s", "12m45s"}

	var timeDurations []time.Duration
	for _, ts := range timeStrings {
		d, err := time.ParseDuration(ts)

		if err != nil {
			log.Fatalf("Time Duration parse failed: %v", err)
		}

		timeDurations = append(timeDurations, d)
	}

	m.Downloads = []Download{
		{
			Id:            0,
			Title:         "Download 0",
			Url:           "https://youtube.com/0",
			Status:        COMPLETED,
			DownloadPath:  "/downloads/one",
			Progress:      100,
			TimeRemaining: timeDurations[0],
		},
		{
			Id:            1,
			Title:         "Download 1",
			Url:           "https://youtube.com/1",
			Status:        ERROR,
			DownloadPath:  "/downloads/two",
			Progress:      67,
			TimeRemaining: timeDurations[1],
		},
		{
			Id:            2,
			Title:         "Download 2",
			Url:           "https://youtube.com/2",
			Status:        IN_PROGRESS,
			DownloadPath:  "/downloads/three",
			Progress:      50,
			TimeRemaining: timeDurations[2],
		},
		{
			Id:            3,
			Title:         "Download 3",
			Url:           "https://youtube.com/3",
			Status:        PENDING,
			DownloadPath:  "/downloads/four",
			Progress:      0,
			TimeRemaining: timeDurations[3],
		},
		{
			Id:            3,
			Title:         "Dragonflight Crests Got Uncapped... But They're Still Not Account Wide",
			Url:           " https://www.youtube.com/watch?v=nL7XxVZOqeg",
			Status:        PENDING,
			DownloadPath:  "/downloads/four",
			Progress:      0,
			TimeRemaining: timeDurations[3],
		},
	}
	m.NextId = 5
}

func (m *Mock) NewDownload(title string, url string, downloadPath string) Download {
	d := Download{
		Id:           m.NextId,
		Title:        title,
		Url:          url,
		DownloadPath: downloadPath,
		Status:       PENDING,
		Progress:     0,
	}

	m.Downloads = append(m.Downloads, d)
	m.NextId++
	return d
}
