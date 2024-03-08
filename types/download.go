package types

import "time"

type Download struct {
	Id            int64
	Title         string
	DownloadPath  string
	Url           string
	Status        Status
	TimeRemaining time.Duration
	Progress      float32
}
