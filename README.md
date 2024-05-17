![go workflow](https://github.com/fmiskovic/new-amazon/actions/workflows/go-ci.yml/badge.svg)
![lint workflow](https://github.com/fmiskovic/new-amazon/actions/workflows/golangci-lint.yml/badge.svg)

## Online Book Store Example

This project is an example of an Online Book Store. It's designed to demonstrate how to create a backend API using Go programming language and a Postgres database to store book and order information.

### Getting Started

These instructions will guide you through getting a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Before you begin, ensure you have Docker installed on your machine. Docker is required to run the Postgres database in a container.

### How to run

Follow these steps to get your development environment running:

1. Clone the repository

2. From terminal, navigate to the project directory

3. Build the application:

```bash
make build
```
4. Start the Postgres database:

```bash
make run-db
```

 - This command starts a PostgreSQL container in the background.

5. Init the database migration tables, `make` command:

```bash
make db cmd=init
```

6. Migrate and seed the database:

```bash
make db cmd=migrate
```

7. Start the application:

```bash
make run
```

### Swagger Documentation

The Swagger documentation for the API is available at [http://localhost:8080/docs](http://localhost:8080/docs) once the application is running. This documentation provides a detailed overview of the available API endpoints, their parameters, and responses.

### Commands

Visit the `Makefile` in the root dir to see all available commands.

## Built With

* [Go](https://golang.org/) - The programming language used.
* [Bun](https://bun.uptrace.dev/guide/) - ORM for Go.
* [Echo](https://echo.labstack.com/) - Web framework for Go.
* [PostgreSQL](https://www.postgresql.org/) - Database Management.
* [Docker](https://www.docker.com/) - Containerization.
