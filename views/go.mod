module media-download-manager/views

go 1.22.1

require media-download-manager/types v0.0.0-00010101000000-000000000000

require (
	github.com/maragudk/gomponents v0.20.3
	media-download-manager/views/components v0.0.0-00010101000000-000000000000
)

replace media-download-manager/types => ../types

replace media-download-manager/views/components => ./components
