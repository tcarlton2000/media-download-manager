module main

go 1.22.1

replace media-download-manager/modules => ./modules

replace media-download-manager/app => ./app

require media-download-manager/app v0.0.0-00010101000000-000000000000

require (
	github.com/maragudk/gomponents v0.20.3 // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	media-download-manager/db v0.0.0-00010101000000-000000000000 // indirect
	media-download-manager/modules v0.0.0-00010101000000-000000000000 // indirect
	media-download-manager/types v0.0.0-00010101000000-000000000000 // indirect
	media-download-manager/views v0.0.0-00010101000000-000000000000 // indirect
	media-download-manager/views/components v0.0.0-00010101000000-000000000000 // indirect
)

replace media-download-manager/db => ./db

replace media-download-manager/types => ./types

replace media-download-manager/views => ./views

replace media-download-manager/views/components => ./views/components
