# FIBER API

The project is an API built in Go using the Fiber framework.

## Installation

1. Fill in the environment variables of the `.template.env` file and rename it to `.env`.
2. Build the project database using: 
```bash
docker-compose up -d
```
3. Install project dependencies using: 
```bash
go mod download
```
4. Run the project using:
```bash
go run cmd/main.go

# or

air
```

## Instructions

All the routes are documented in the path `/swagger`.

> Every time you make changes in the api and you want to rebuild the documentation you must use this command, in the root of the project:
> ```bash
>   swag init --pd
> ```

## Technologies

* [Fiber](https://gofiber.io/)
* [Gorm](https://gorm.io/)
* [Swagger](https://github.com/gofiber/swagger)