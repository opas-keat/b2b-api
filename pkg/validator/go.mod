module validator

go 1.19

require (
	github.com/go-playground/locales v0.14.0
	github.com/go-playground/universal-translator v0.18.0
	github.com/go-playground/validator/v10 v10.11.1
	models v0.0.0
	shareerrors v0.0.0
)

require (
	github.com/leodido/go-urn v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3 // indirect
	golang.org/x/sys v0.0.0-20210806184541-e5e7981a1069 // indirect
	golang.org/x/text v0.3.7 // indirect
)

replace (
	models v0.0.0 => ../models
	shareerrors v0.0.0 => ../share-errors
)
