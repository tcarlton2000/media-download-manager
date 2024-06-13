package db

import (
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
	readDb  *ReadDb
	writeDb *WriteDb
}

func OpenDb(configDir string) *Database {
	writeDb := OpenWriteDb(configDir)
	readDb := OpenReadDb(configDir)

	return &Database{readDb: readDb, writeDb: writeDb}
}

func (db *Database) GetDownloads() ([]types.Download, error) {
	return db.readDb.GetDownloads()
}

func (db *Database) NewDownload(d types.Download) (int64, error) {
	return db.writeDb.NewDownload(d)
}

func (db *Database) UpdateDownload(d types.Download) error {
	return db.writeDb.UpdateDownload(d)
}

func (db *Database) DeleteDownload(id int64) error {
	return db.writeDb.DeleteDownload(id)
}
