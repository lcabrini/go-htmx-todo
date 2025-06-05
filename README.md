# go-htmx-todo

My own contribution to the world of basic Todo-list apps. I am using:
1. [Go](https://go.dev)
1. http.ServeMux for routing
1. [PostgreSQL](https://www.postgresql.org/)
1. [HTMX](https://htmx.org/)
1. [sqlc](https://sqlc.dev/)
1. [Pico CSS](https://picocss.com/)
1. [Goose](https://pressly.github.io/goose/)

## Setting up

Make sure you have goose installed:

``` console
$ go install github.com/pressly/goose/v3/cmd/goose@latest
```

If you want to poke at the code, you will also want to have sqlc installed:

``` console
$ go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

Create `.env`, with the following environment variables. Adjust for your postgresql setup.

``` bash
PORT=8000
DB_URL=postgres://user:password@localhost:5432/dbname?sslmode=disable
```
Run the migrations:

``` console
$ cd sql/schema
$ goose postgres postgres://user:password@localhost:5432/dbname?sslmode=disable
$ cd ../..
```

Build the code:

``` console
$ go mod tidy
$ go build
```

Run the server:

``` console
$ ./go-htmx-todo
```

Browse to `https://localhost:8000` (or whatever port you specified in the `.env` file)

