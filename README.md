# chat-server

## goenv

https://github.com/syndbg/goenv

## direnv(PATH)

https://github.com/direnv/direnv

## dep（module management）

https://github.com/golang/dep

## echo（framework）

https://github.com/labstack/echo

## usage

```
install goenv
公式を参照

install direnv
公式を参照

move to project dir
$ cd src/app

install modules
$ dep ensure -update

server start
$ go run *.go

format code
$ go fmt

test
$ go test
```

## api

```
login
$ curl -X POST -d 'username=test' -d 'password=test' localhost:1323/login
```
