package db

import (
	"database/sql"
	"log"

	"media-download-manager/types"

	_ "github.com/mattn/go-sqlite3"
)

const create string = `
	CREATE TABLE IF NOT EXISTS downloads (
		id INTEGER NOT NULL PRIMARY KEY,
		title TEXT NOT NULL,
		url TEXT NOT NULL,
		download_path TEXT NOT NULL,
		status INTEGER NOT NULL,
		progress REAL NOT NULL,
		time_remaining_ms INTEGER NOT NULL
	);
`

type Database struct {
	db *sql.DB
}

func OpenDb() *Database {
	db, err := sql.Open("sqlite3", "./media-download-manager.db")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(create); err != nil {
		log.Fatal(err)
	}

	return &Database{db: db}
}

func (db *Database) GetDownloads() ([]types.Download, error) {
	rows, err := db.db.Query("SELECT * FROM downloads")
	if err != nil {
		log.Print(err)
		return []types.Download{}, err
	}
	defer rows.Close()

	var downloads []types.Download
	for rows.Next() {
		var d types.Download
		err = rows.Scan(&d.Id, &d.Title, &d.Url, &d.DownloadPath, &d.Status, &d.Progress, &d.TimeRemaining)
		if err != nil {
			log.Print(err)
			return []types.Download{}, err
		}

		downloads = append(downloads, d)
	}

	return downloads, nil
}

func (db *Database) NewDownload(d types.Download) (int64, error) {
	stmt, err := db.db.Prepare("INSERT INTO downloads(title, url, download_path, status, progress, time_remaining_ms) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Print(err)
		return 0, err
	}

	res, err := stmt.Exec(d.Title, d.Url, d.DownloadPath, d.Status, d.Progress, d.TimeRemaining)
	if err != nil {
		log.Print(err)
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Print(err)
		return 0, err
	}

	return id, nil
}

func (db *Database) DeleteDownload(id int64) error {
	stmt, err := db.db.Prepare("DELETE FROM downloads WHERE id = ?")
	if err != nil {
		log.Print(err)
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}
