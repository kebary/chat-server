# chat-server

## goenv

https://github.com/syndbg/goenv

## direnv(PATH)

https://github.com/direnv/direnv

## dep（module management）

https://github.com/golang/dep

## echo（framework）

https://github.com/labstack/echo

## ORM

https://github.com/jinzhu/gorm

## live reload

https://github.com/pilu/fresh

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
$ fresh
（command not foundと言われた下記を参照してPATH設定を）
https://github.com/pilu/fresh/issues/74

format code
$ go fmt ./...

test
$ go test
```

## api

```
/login
$ curl -X POST -H 'Content-Type: application/json' -d '{"username":"test","password":"test"}' localhost:1323/login

/talk
$ curl -X GET localhost:1323/api/v1/talk -H "Authorization: Bearer token"
$ curl -X POST -H 'Content-Type: application/json' -H "Authorization: Bearer token" -d '{"msg":"curl test"}' localhost:1323/api/v1/talk
```
