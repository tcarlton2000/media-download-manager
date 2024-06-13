package db

import (
	"database/sql"
	"fmt"
	"log"
	"media-download-manager/types"
	"sync"
	"time"
)

type ReadDb struct {
	sync.Mutex
	db *sql.DB
}

func OpenReadDb(configDir string) *ReadDb {
	readDb, err := sql.Open("sqlite3", fmt.Sprintf("file:%s/media-download-manager.db?mode=ro", configDir))
	if err != nil {
		panic(err)
	}

	return &ReadDb{db: readDb}
}

func (rdb *ReadDb) GetDownloads() ([]types.Download, error) {
	rows, err := rdb.db.Query("SELECT * FROM downloads ORDER BY id DESC")
	if err != nil {
		log.Print(err)
		return []types.Download{}, err
	}
	defer rows.Close()

	var downloads []types.Download
	for rows.Next() {
		var d types.Download
		var timeRemainingMs int64
		err = rows.Scan(&d.Id, &d.Title, &d.Url, &d.DownloadPath, &d.Status, &d.Progress, &timeRemainingMs)
		if err != nil {
			log.Print(err)
			return []types.Download{}, err
		}

		d.TimeRemaining = time.Duration(timeRemainingMs) * time.Millisecond

		downloads = append(downloads, d)
	}

	return downloads, nil
}
