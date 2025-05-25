<img src="image.gif">

## introduce

Only once click, automatically activated

## dev

install go-bindata

```bash
go install github.com/go-bindata/go-bindata/v3/go-bindata@latest
export PATH=$PATH:$(go env GOPATH)/bin
go-bindata --version
go-bindata -o internal/util/access.go -pkg util static/... templates/...
go run cmd/main.go
```

## run it !

mac linux windows

```
make run
```

## make it !

mac or linux ：

```bash
make run
make build-all
make clean
```

windows use powershell run:

```powershell
.\build.ps1
```

## Stargazers over time

[![Stargazers over time](https://starchart.cc/saxpjexck/lsix.svg?variant=adaptive)](https://starchart.cc/saxpjexck/lsix)
