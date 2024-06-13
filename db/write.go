package db

import (
	"database/sql"
	"fmt"
	"log"
	"media-download-manager/types"
	"sync"
)

type WriteDb struct {
	sync.Mutex
	db *sql.DB
}

func OpenWriteDb(configDir string) *WriteDb {
	writeDb, err := sql.Open("sqlite3", fmt.Sprintf("file:%s/media-download-manager.db?mode=rwc&_journal_mode=WAL", configDir))
	if err != nil {
		panic(err)
	}

	// SQLite can have only 1 open write connection at a time.
	writeDb.SetMaxOpenConns(1)
	if _, err := writeDb.Exec(create); err != nil {
		panic(err)
	}

	return &WriteDb{db: writeDb}
}

func (wdb *WriteDb) NewDownload(d types.Download) (int64, error) {
	wdb.Lock()
	stmt, err := wdb.db.Prepare("INSERT INTO downloads(title, url, download_path, status, progress, time_remaining_ms) values (?, ?, ?, ?, ?, ?)")
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

	wdb.Unlock()
	return id, nil
}

func (wdb *WriteDb) UpdateDownload(d types.Download) error {
	wdb.Lock()
	stmt, err := wdb.db.Prepare("UPDATE downloads SET title=?, status=?, progress=?, time_remaining_ms=? WHERE id=?")
	if err != nil {
		log.Print(err)
		wdb.Unlock()
		return err
	}

	_, err = stmt.Exec(d.Title, d.Status, d.Progress, d.TimeRemaining.Milliseconds(), d.Id)
	if err != nil {
		log.Print(err)
		wdb.Unlock()
		return err
	}

	wdb.Unlock()
	return nil
}

func (wdb *WriteDb) DeleteDownload(id int64) error {
	wdb.Lock()
	stmt, err := wdb.db.Prepare("DELETE FROM downloads WHERE id = ?")
	if err != nil {
		log.Print(err)
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Print(err)
		return err
	}

	wdb.Unlock()
	return nil
}
