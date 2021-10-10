# outerbanks-api
A graphql api where we can store and get information on characters in Outerbanks.

Fetch library using `go get`
https://github.com/99designs/gqlgen

```go get github.com/99designs/gqlgen```

Add gqlgen to your project's tools.go
```
printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > tools.go
go mod tidy
```

Start the graphql server
```
go run server.go
```


