package db

import (
	"database/sql"
	"log"
	"sync"
	"time"

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
	sync.Mutex
	readDb  *sql.DB
	writeDb *sql.DB
}

func OpenDb() *Database {
	writeDb, err := sql.Open("sqlite3", "file:media-download-manager.db?mode=rw")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := writeDb.Exec(create); err != nil {
		log.Fatal(err)
	}

	readDb, err := sql.Open("sqlite3", "file:media-download-manager.db?mode=ro")
	if err != nil {
		log.Fatal(err)
	}

	return &Database{readDb: readDb, writeDb: writeDb}
}

func (db *Database) GetDownloads() ([]types.Download, error) {
	rows, err := db.readDb.Query("SELECT * FROM downloads ORDER BY id DESC")
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

func (db *Database) NewDownload(d types.Download) (int64, error) {
	db.Lock()
	stmt, err := db.writeDb.Prepare("INSERT INTO downloads(title, url, download_path, status, progress, time_remaining_ms) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Print(err)
		return 0, err
	}

	res, err := stmt.Exec(d.Title, d.Url, d.DownloadPath, d.Status, d.Progress, d.TimeRemaining.Milliseconds())
	if err != nil {
		log.Print(err)
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Print(err)
		return 0, err
	}

	db.Unlock()
	return id, nil
}

func (db *Database) UpdateDownload(d types.Download) error {
	db.Lock()
	stmt, err := db.writeDb.Prepare("UPDATE downloads SET title=?, status=?, progress=?, time_remaining_ms=? WHERE id=?")
	if err != nil {
		log.Print(err)
		db.Unlock()
		return err
	}

	log.Printf("Executing %f progress update", d.Progress)
	_, err = stmt.Exec(d.Title, d.Status, d.Progress, d.TimeRemaining.Milliseconds(), d.Id)
	if err != nil {
		log.Print(err)
		db.Unlock()
		return err
	}

	db.Unlock()
	return nil
}

func (db *Database) DeleteDownload(id int64) error {
	db.Lock()
	stmt, err := db.writeDb.Prepare("DELETE FROM downloads WHERE id = ?")
	if err != nil {
		log.Print(err)
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Print(err)
		return err
	}

	db.Unlock()
	return nil
}
