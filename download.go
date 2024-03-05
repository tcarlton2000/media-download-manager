package main

import "time"

type Download struct {
	Id            int
	Title         string
	Url           string
	Status        Status
	TimeRemaining time.Duration
	Progress      float32
}

type DownloadRow struct {
	Download      Download
	ProgressProps ProgressProps
}

func (d Download) DownloadRow() DownloadRow {
	return DownloadRow{Download: d, ProgressProps: ProgressProps{Progress: d.Progress}}
}
