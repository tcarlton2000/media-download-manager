module media-download-manager/modules

go 1.22.1

require (
	media-download-manager/db v0.0.0-00010101000000-000000000000
	media-download-manager/types v0.0.0-00010101000000-000000000000
)

require github.com/mattn/go-sqlite3 v1.14.22 // indirect

replace media-download-manager/types => ../types

replace media-download-manager/db => ../db
