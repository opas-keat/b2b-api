module users

go 1.19

require (
	github.com/caarlos0/env/v6 v6.10.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/stretchr/testify v1.8.0
	go.uber.org/automaxprocs v1.5.1
	go.uber.org/dig v1.15.0
	gorm.io/gorm v1.24.0
	models v0.0.0
)

require (
	github.com/AlekSi/pointer v1.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.13.0 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.12.0 // indirect
	github.com/jackc/pgx/v4 v4.17.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.4.0 // indirect
	golang.org/x/crypto v0.0.0-20220722155217-630584e8d5aa // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	gorm.io/driver/postgres v1.4.4
)

replace models => ../../pkg/models
