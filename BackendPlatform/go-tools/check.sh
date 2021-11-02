# check code
golangci-lint run

# test
go test -race -cover $(go list ./... | grep -v "vendor")
