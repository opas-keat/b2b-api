module user

go 1.19

require (
	fibercore v0.0.0
	github.com/caarlos0/env/v6 v6.10.1
	github.com/gofiber/fiber/v2 v2.40.1
	github.com/tidwall/gjson v1.14.4
	go.uber.org/automaxprocs v1.5.1
	go.uber.org/dig v1.15.0
	models v0.0.0
)

require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.41.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab // indirect
)

replace (
	fibercore v0.0.0 => ../../pkg/fiber-core
	models v0.0.0 => ../../pkg/models
)
