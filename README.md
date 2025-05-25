<img src="image.gif">

## introduce

**Only once click, automatically activated**

- Click the ide card to automatically inject ja-netfilter and activation key, and restart the IDE to see that it has been activated.
- Click the plugin or ide cards to copy the activation key to the clipboard, and you can manually enter the key for activation
- It also supports one-click removal of activation configuration
- It is single bin file, after execution, the relevant files for activation will be released in the current directory. After activation, the ide will reference the directory and do not delete it.

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

## Star History

[![Stargazers over time](https://starchart.cc/saxpjexck/lsix.svg?variant=adaptive)](https://starchart.cc/saxpjexck/lsix)
