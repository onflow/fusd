
To generate the `internal/assets` directory:

```
# install go-bindata
go get -u github.com/kevinburke/go-bindata/...

# generate for contracts
go-bindata -o internal/assets/assets.go ../../../contracts
```