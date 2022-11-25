# Movies Library
A movies library that gives the users ability to add movies to the library and save their watched movies.

Written using Domain Driven Design and clean architecture principles.

Written with [golang](https://github.com/golang/go), [gorm](https://github.com/go-gorm/gorm) and [gin](https://github.com/gin-gonic/gin)

# Get Started
`cp .env.example .env`

`docker-compose build`

`docker-compose up`

`go run .`

# API
[Collection](https://elements.getpostman.com/redirect?entityId=130550-351a9b84-d08e-472a-b466-9e5789b2e0db&entityType=collection), [Documentation](https://documenter.getpostman.com/view/130550/2s8YswQBnZ)

# Architecture
```
.
├── app
│   ├── controllers
│   │   ├── auth.go
│   │   ├── helpers.go
│   │   ├── movie.go
│   │   ├── user.go
│   │   └── watched_list.go
│   ├── forms
│   │   ├── auth.go
│   │   ├── movie.go
│   │   ├── user.go
│   │   └── watched_list.go
│   ├── middleware
│   │   └── auth.go
│   └── router.go
├── database
│   └── migrations
│       └── schema.sql
├── docker
│   └── mysql
│       ├── conf.d
│       │   └── my.cnf
│       └── log
│           └── mysqld_general.log
├── docker-compose.yml
├── domain
│   ├── entities
│   │   ├── movie.go
│   │   ├── user.go
│   │   ├── user_token.go
│   │   └── watched_list.go
│   ├── repositories
│   │   ├── movie.go
│   │   ├── user.go
│   │   ├── user_token.go
│   │   └── watched_list.go
│   └── services
│       ├── auth.go
│       ├── encryptor.go
│       ├── movie.go
│       ├── user.go
│       └── watched_list.go
├── go.mod
├── go.sum
├── infrastructure
│   ├── init.go
│   ├── repositories
│   │   ├── movie.go
│   │   ├── user.go
│   │   ├── user_token.go
│   │   └── watched_list.go
│   └── services
│       └── encryptor.go
├── main.go
└── README.md

17 directories, 39 files
```

# Database Diagram
<img src="https://github.com/HazemNoor/movies-library/raw/main/db-giagram.png?raw=true" alt="Database Diagram" width="500">