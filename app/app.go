package app

import "media-download-manager/modules"

type App struct {
	mock modules.Mock
}

func (a *App) Init() {
	a.mock.Init()
}