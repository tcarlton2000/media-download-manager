module media-download-manager/modules

go 1.22.1

require (
	github.com/kkdai/youtube/v2 v2.10.0
	media-download-manager/db v0.0.0-00010101000000-000000000000
	media-download-manager/types v0.0.0-00010101000000-000000000000
)

require (
	github.com/bitly/go-simplejson v0.5.1 // indirect
	github.com/dlclark/regexp2 v1.10.0 // indirect
	github.com/dop251/goja v0.0.0-20231027120936-b396bb4c349d // indirect
	github.com/go-sourcemap/sourcemap v2.1.3+incompatible // indirect
	github.com/google/pprof v0.0.0-20231101202521-4ca4178f5c7a // indirect
	github.com/mattn/go-sqlite3 v1.14.22 // indirect
	golang.org/x/text v0.14.0 // indirect
)

replace media-download-manager/types => ../types

replace media-download-manager/db => ../db
