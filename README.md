# Testing Basics

## Commands

```
go test
go test -v
go test *.go
go test *.go -tags=integration -v
go test *.go -tags=unit -v 
go test *.go -cover
```

To view test coverage
```
go test *.go -coverprofile=c.out
go tool cover -html=c.out
```