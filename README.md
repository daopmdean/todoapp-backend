# todoapp-backend

Rest API for todoapp written in [Golang](https://golang.org/), using Clean Architecture and CI/CD (includes unit tests and integration tests).

Using:

- Web framework: [Gin](https://github.com/gin-gonic/gin)
- Authentication: [JWT](https://jwt.io/)
- Database (choose one out of three options):
  - In Memory DB
  - Microsoft SQL Server (using [Gorm](https://gorm.io/docs/))
  - MongoDB (coming soon)
- CI/CD:
  - GitHub Actions (Docker Image for testing)
  - Microsoft Azure Web App (Docker Image for deployment)

> Note: Config to change database (default to In Memory DB)

## Getting Started

Development environment needed:

- [Golang](https://golang.org/dl/) (>v1.13)

Clone the project to your working directory

```zsh
git clone https://github.com/daopmdean/todoapp-backend.git
```

This project support to run on

- Memory Database
- Microsoft SQL Server Database

Setup env on zsh

```zsh
export <ENV_NAME>='<ENV_VALUE>'
```

check to see if your environment declared

```zsh
echo $<ENV_NAME>
```

Example

```zsh
export NAME='Dean'
echo $NAME
-> Dean
```

PORT: environment variable to setup the port your program running on.<br>
If you don't set PORT env variable, Program will run on default port: http://localhost:5000<br>
Example:

```zsh
export PORT=3456
Your program will run on -> http://localhost:3456
```

## DB Mode

You can choose DB Mode to run with DB_MODE environment variable. If you do not provide any, default the program will run with memory db.

### Memory DB

Run the program

```zsh
make run
```

### MS SQL DB

```zsh
export DB_MODE=mssql
```

Using ORM Library: [Gorm](https://gorm.io/index.html)<br>
Setup env variables

- DB_USERNAME: Username to connect to sql server
- DB_PASSWORD: Password to connect to sql server
- DB_HOST: Host to connect to sql server
- DB_PORT: Port to connect to sql server
- DB_NAME: Database name to connect to sql server<br>

Example

```zsh
export DB_USERNAME=sa
export DB_PASSWORD=MyPassword@123
export DB_HOST=localhost
export DB_PORT=1433
export DB_NAME=todoapp
```

Run the program

```zsh
make run
```

## Test

To run all tests

```zsh
make test_all
```

To run api tests

```zsh
make test_api
```

To run usecase tests

```zsh
make test_usecase
```

To run entity tests

```zsh
make test_entity
```

## System Architect

The system designed following [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) guideline.

![Clean Architect](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

### Project structure

```
todoapp
├── cmd/server
│   └── main.go
├── internal
│   └── adapter
│       └── http/rest
│       └── memdb
│       └── mssqldb
│   └── common/config
│   └── entity
│   └── usecase
│       └── repo
└── test
    └── api_rest_test
    └── repomock
```

### Annotate

- cmd/server/main.go: run the program
- adapter: store plugins that you can easily replace with another implementations
- http/rest: store api handle http request with gin-gonic framework
- memdb: implementations for memory db
- mssqldb: implementations for microsoft sql server db
- test: store test cases
