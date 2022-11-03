# SMPLRSTAPP [Backend application]
This simple and test rest api project on go lang.
_______________________________

We try create application on clean architecture on go lang and use best practices:
    https://github.com/golang-standards/project-layout

It's contain:
    postgresql as db
    gorm as orm

## Build & Run (Locally)
### Prerequisites
- go 1.19
- docker & docker-compose
- [golangci-lint](https://github.com/golangci/golangci-lint) (<i>optional</i>, used to run code checks)
- [swag](https://github.com/swaggo/swag) (<i>optional</i>, used to re-generate swagger documentation)
- [gorm](https://github.com/go-gorm/gorm) (<i>optional</i>, used to connect and work with database as ORM)
- [mux](https://github.com/gorilla/mux) (<i>optional</i>, used to HTTP request multiplexer)
- [dotenv](https://github.com/joho/godotenv) (<i>optional</i>, used to load environment of config)

Create .env file in root directory and add following values:
```dotenv
APP_ENV=local

POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USERNAME=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DATABASE=smplrstapp

HTTP_HOST=localhost
HTTP_PORT=8020

Use `make run` to build&run project, `make lint` to check code with linter.