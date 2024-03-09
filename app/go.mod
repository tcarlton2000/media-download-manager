module media-download-manager/app

go 1.22.1

replace media-download-manager/db => ../db

replace media-download-manager/modules => ../modules

replace media-download-manager/types => ../types

require (
	media-download-manager/db v0.0.0-00010101000000-000000000000
	media-download-manager/modules v0.0.0-00010101000000-000000000000
	media-download-manager/types v0.0.0-00010101000000-000000000000
)

require (
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	github.com/wader/goutubedl v0.0.0-20240306161536-c309f999af46 // indirect
)
