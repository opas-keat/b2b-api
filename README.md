f# b2b-api
# Quick start
## Run local all service

    ./run_local.sh

## Run local single service

    cd service services/{service_name} && make local

## Install dependencies
### 1. mockery

    brew install mockery
    brew upgrade mockery

or

    go install github.com/vektra/mockery/v2@latest

### 2. air

    go install github.com/cosmtrek/air@latest

> **Warning**
> If run program then the incident `too many open files` to follow [Here](https://accurate-adapter-c2f.notion.site/Air-de80562ecdd5419c85804d09af8518eb)

### 3. gosec

    go install github.com/securego/gosec/v2/cmd/gosec@latest

> **Warning**
> Please run ``make sec-scan`` before push your feature branch !!!!
# Example mockery
1. write code `//go:generate mockery --name {name_interface}` on top interface
2. run command `go generate ./...`

```go
//go:generate mockery --name Example
type Example interface {
    Print()
    GetString() string
    SetString(s string)
}
```