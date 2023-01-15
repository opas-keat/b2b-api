module logs

go 1.19

require (
	fibercore v0.0.0
	github.com/AlekSi/pointer v1.2.0
	github.com/caarlos0/env/v6 v6.10.1
	github.com/gofiber/fiber/v2 v2.41.0
	github.com/google/uuid v1.3.0
	github.com/jinzhu/copier v0.3.5
	github.com/samber/lo v1.37.0
	go.uber.org/dig v1.16.1
	golang.org/x/exp v0.0.0-20230113213754-f9f960f08ad4
	gorm.io/driver/postgres v1.4.6
	gorm.io/gorm v1.24.3
	models v0.0.0
	shareerrors v0.0.0
	validator v0.0.0
)

require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/cosmtrek/air v1.41.0 // indirect
	github.com/creack/pty v1.1.18 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.11.1 // indirect
	github.com/imdario/mergo v0.3.13 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.2.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.15.14 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.4.3 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.44.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/crypto v0.4.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	fibercore v0.0.0 => ../../pkg/fiber-core
	models v0.0.0 => ../../pkg/models
	shareerrors v0.0.0 => ../../pkg/share-errors
	validator v0.0.0 => ../../pkg/validator
)
