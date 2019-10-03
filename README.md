docker run -it --rm -v $(pwd):/app -w /app golang:1.13-rc-alpine sh -c 'go build -o first_go first_go.go && ./first_go'

References:

[Go by examples](https://gobyexample.com/)

[Learn Golang](https://github.com/golang/go/wiki/Learn)