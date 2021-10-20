
To generate the `internal/assets` directory, run `go generate`.

This will read your source code (`contracts.go`) and run the command:
```
//go:generate go run github.com/kevinburke/go-bindata/go-bindata -prefix ../../../contracts -o internal/assets/assets.go -pkg assets -nometadata -nomemcopy ../../../contracts
```