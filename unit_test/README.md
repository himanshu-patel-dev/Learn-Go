# Unit Testing

Run all test in directory and sub directories
```
$ go test
```

Get coverage of tests
```
# first step
$ go test -coverprofile=coverage.out
# second step
$ go tool cover -html=coverage.out 
```
