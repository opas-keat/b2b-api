module systemlink

go 1.19

require (
	fibercore v0.0.0
	github.com/AlekSi/pointer v1.2.0
	github.com/caarlos0/env/v6 v6.10.1
	github.com/gofiber/fiber/v2 v2.40.1
	github.com/samber/lo v1.37.0
	go.uber.org/dig v1.15.0
	golang.org/x/exp v0.0.0-20221227203929-1b447090c38c
	gorm.io/driver/sqlserver v1.4.1
	gorm.io/gorm v1.24.2
	models v0.0.0
	shareerrors v0.0.0
	validator v0.0.0
)

require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.11.1 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/microsoft/go-mssqldb v0.17.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/stretchr/testify v1.8.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.41.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/crypto v0.0.0-20221005025214-4161e89ecf1b // indirect
	golang.org/x/sys v0.1.0 // indirect
	golang.org/x/text v0.3.7 // indirect
)

replace (
	fibercore v0.0.0 => ../../pkg/fiber-core
	models v0.0.0 => ../../pkg/models
	shareerrors v0.0.0 => ../../pkg/share-errors
	validator v0.0.0 => ../../pkg/validator
)
