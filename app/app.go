package app

import (
	"media-download-manager/db"
)

type App struct {
	db *db.Database
}

func (a *App) Init() {
	a.db = db.OpenDb()
}
