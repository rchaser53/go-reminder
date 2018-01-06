# memo

i used [aws-lambda-go-shim](https://github.com/eawsy/aws-lambda-go-shim).
the below is the how to setup

```
$ docker pull eawsy/aws-lambda-go-shim:latest
$ go get -u -d github.com/eawsy/aws-lambda-go-core/...
$ wget -O Makefile https://git.io/vytH8
$ make
```