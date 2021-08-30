# Golang OpenAPI Server Template
A backend API server template based on Gin, with DB connection of PostgreSQL, Redis and MongoDB.

This server is built using [Golang](https://golang.org/) as programming language and [Gin](https://gin-gonic.com/) as web framework, following OpenAPI 3.0 Specification.

[GoFiber](https://gofiber.io/) version can be found in [this branch](https://github.com/KOVERcjm/go-openapi-server-template/tree/GoFiber_based).

# 0 Getting started
Install dependency and initialize DB via Docker (if not existed):

``` shell
go mod tidy

# Generate self-signed cert if necessary.
openssl req -x509 -new -nodes -sha256 -utf8 -days 365 -newkey rsa:4096 -subj '/CN=localhost' -keyout cert/private.key -out cert/certificate.crt

# Start DB via Docker if necessary
docker run --name PostgreSQL --restart always -e POSTGRES_PASSWORD=12345 -e POSTGRES_DB=examplePg -p 5432:5432 -d postgres

docker run --name Mongo --restart always -e MONGO_INITDB_ROOT_USERNAME=mongo -e MONGO_INITDB_ROOT_PASSWORD=12345 -e MONGO_INITDB_DATABASE=exampleMongo -p 27017:27017 -d mongo

docker run --name Redis --restart always -p 6379:6379 -d redis redis-server --appendonly yes --requirepass "12345"
```

Test server (TODO):

```shell
go test
```

Run server in development:

```shell
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
$(go env GOPATH)/bin/air
```

Run server in production (compile and then start):

```shell
go build -o cmd/app kovercheng
cmd/app
```


# 2 How to use

1. Search and replace all 'kovercheng' in this project with your project name.
2. Replace all '3000' in `.env Dockerfile` with actual port number you want to use.
3. Have fun.


# 3 Project Directories

- `/cert` - place the server's HTTPS crt and key file here (remind to use the HTTPS codes in `main.go`)
- `/cmd` - compiled application executables
- `/driver` - establish connection between handlers and external drivers (mainly database drivers now)
- `/handler` - api server handler for business logics
- `/middlewares` - store go-fiber middlewares like logger
- `/model` - store structs for database and business handler use
  - `/document` - Mongo document structs
  - `/table` - Postgres table structs
- `service` - internal helper for handlers to use, like DB service


# 4 Dockerize

Replace the DB connection string in Dockerfile if not have DB server in local.

Run `docker build` command to pack the server into an image. The following is for your reference:

``` shell
docker build -t my_app .
docker run -d --name my_app -p 3000:3000 my_app
```


# 5 Reference

- [OpenAPI 3.0 Specification](https://swagger.io/specification/)
- [standard-go-project-layout](https://github.com/golang-standards/project-layout)
- [gp-fiber](https://github.com/gin-gonic/gin)
