# REST API with Golang and MongoDB Demo

This repository contains codes of my golang restapi demo on mongodb.


1. Install mongodb on your local computer.

2. Run following commands on the terminal.
```
go get gopkg.in/mgo.v2/bson

go get go.mongodb.org/mongo-driver/mongo

go get go.mongodb.org/mongo-driver/mongo/options

go get github.com/gin-gonic/gin

go mod init crud-api

go mod tidy

go run main.go
```

It runs on port 8080 by default.




1. GET All Users:
get http://localhost:8080/users

```
[
  {
    "id": "363338653138633132646337353735653333623163653163",
    "name": "asd",
    "address": "XY Str. 12345",
    "age": 25,
    "gender": "xxxxxxx"
  },
  {
    "id": "363338653139363736396136366631343731326336363032",
    "name": "aaaaaaa",
    "address": "XY Str. 12345",
    "age": 25,
    "gender": "xxxxxxx"
  },
  {
    "id": "363338653139636132303666386561623063316137306263",
    "name": "aaaaaaa",
    "address": "XY Str. 12345",
    "age": 25,
    "gender": "xxxxxxx"
  }
]
```


2. Create a User
POST http://localhost:8080/users/create

```
{
  "name" : "aaaaaaa",
	"address" : "XY Str. 12345",
	"age" : 25,
	"gender":"xxxxxxx"
}
```

Response:
```
{
  "Value Inserted": {
    "id": "",
    "name": "aaaaaaa",
    "address": "XY Str. 12345",
    "age": 25,
    "gender": "xxxxxxx"
  }
}
```

3. Delete a User:
DELETE http://localhost:8080/users/asd/delete

```
{
  "Deleted User": {
    "id": "363338653138633132646337353735653333623163653163",
    "name": "asd",
    "address": "XY Str. 12345",
    "age": 25,
    "gender": "xxxxxxx"
  }
}
```

4. Get User by Name
GET http://localhost:8080/users/test

```
[
  {
    "id": "363338653234303964653734336536343865613463313835",
    "name": "test",
    "address": "XY Str. 12345",
    "age": 25,
    "gender": "xxxxxxx"
  }
]
```


