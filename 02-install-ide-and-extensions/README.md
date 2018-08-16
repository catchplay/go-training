#Install IDE and extensions

## Download vscode
https://code.visualstudio.com/download

## Install go extensions
https://code.visualstudio.com/docs/languages/go

### Set up go.gopath in user settings:
```json
{
"go.gopath": "$home/go",
}

```


## Debugging
Debug your code, binaries or tests (using [delve](https://github.com/derekparker/delve))
```
go get -u github.com/derekparker/delve/cmd/dlv
```