# Blueprint: Service Go

This repository contains a blueprint for building backend services/APIs in *Go*.

It is inspired by the [Golang + Gin real-world exmaple application](https://github.com/gothinkster/golang-gin-realworld-example-app).

## Used Libraries

- RESTful JSON API build on top of the [Echo framework](https://echo.labstack.com/)
- [GORM object mapper](http://gorm.io/) for data access (including [support for migrations](https://github.com/go-gormigrate/gormigrate)) with support for
  - SQLite3
  - Postgres
  - MySQL
- Level-based logging using [glg](https://github.com/kpango/glg)
- Env var based configuration using [godotenv](github.com/joho/godotenv)

## Implemented Features

- Configuration through env variables and profile files
- JWT/Token-based authoriazation with OAuth 2.0 workflow for the following grant types:
  - Password
  - Refresh token
- Role-based permission system
- User API with CRUD methods

See the [Postman](https://www.getpostman.com/) collection in `tools/` to test out the API.

## Setup

### Prerequisites

- git
- Go 1.12 or higher

### Start the Application

1. Clone the repo
2. Run `go install` to install all dependencies
3. Run `go run .` to start a development version of the application
4. The API is now accessible under `http://localhost:8080`

## Configuration

The application is fully configurable through environment variables which will be fetched at start-up and converted into a configuration struct.

### Profiles

You can specify active profiles by setting `PROFILES=prof1,prof2`. Thereby the following files will be sourced:

1. `.env.prof1.local`
2. `.env.prof1`
3. `.env.prof2.local`
4. `.env.prof2`

Already present env variables will not be overriden so the order of profiles is important.

The default profile is `development` which runs *Echo* and *GORM* in debug mode. In production include the `production` profile to disable debug modes.

### Variables

See the following table for all supported variables and their meaning.

| Variable     | Description                                             | Default       | Example             |
|--------------|---------------------------------------------------------|---------------|---------------------|
| `PROFILES`   | Active profiles for loading further configuration files | `development` | `production, local` |
| `PUBLIC_URL` | Public base URL                                         | `http://localhost:8080` | `https://api.example.io/myapi` |
| `PORT`       | Local port on which the server listens                  | `8080`        | `443` |
| `DB_TYPE`    | The type of relational database to use, one of `sqlite3`, `postgres`, `mysql`, `mssql` | `sqlite3` | `postgres` |
| `DB_URL`     | DB URL or connection string                             | `database.db` | `"host=myhost port=myport user=gorm dbname=gorm password=mypassword"` |
| `DB_MIGRATE`          | Boolean value indicating whether migrations should be run at start-up | `true` | `false`  |
| `JWT_ALGORITHM`       | JWT algorithm to use, one of `HS256`, `HS384`, `HS512`                | `HS256`| `HS512`  |
| `JWT_KEY`             | The key for encrypting and decrypting JWTs                            | random | `123key` |
| `JWT_EXPIRATION_TIME` | Expiration time for access tokens                                     | `5m`   | `1h`     |
| `JWT_REFRESH_TIME`    | Expiration time for refresh tokens                                    | `72h`  | `60m`    |

## Building

Use `go build` to create a standard dev build.

Run `go build -ldflags "-s -w -X main.release=true -X main.version=<version> -X main.buildDate=<date>"` to create a release build without debug information and with explicit version and build date written into the binary. The `<version>` and `<date>` parameters can be arbitrary string. Sepcifying `main.release=true` will make sure the binary will always be executed with an active `PRODUCTION` profile.

Alternatively, you may want to use the example `makefile` for executing common build tasks.
