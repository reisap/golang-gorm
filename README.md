# Create template clean architecture for golang development inspired by NestJS 

    - Golang 1.22
    - gorm (Mysql)
    - Redis
    - Air (like nodemon for developement)
    - https://dbdiagram.io/ (tools for speed up create ERD and export sql)

## Features
    - Version of api 
    you can separated routing v1,v2,... etc for development
    - dockerize your code
    - Controller, Service, Repository, Entity (model) Yap !!! this separated you can focus which one your code on there, no more spaghetti code :D looks like uncle BOB said
    - global pagination each entity/model
    - compress response
    - helmet for default security
    - cors
    - pprof
    - test with separated design code

## Development Separated into 2 branch
    - main
    use domain model into package inspired by NestJS

    - research 
    use MVC model architecture into package

## Database
    - Default
    can used with migrate from sql query file in folder sql
    make migrateup //command in terminal

    - Code migration from golang
    use from entity in each domain by default in file main.go

## Database Design

![Database Design](./bwacourse.png)

## Run Code
    - make start //docker must be running first
    - make migrate-up //migration database (optional) you can using in code based

## Lint Code
    - GOFLAGS=-buildvcs=false golangci-lint run -v
    debug if any code "Import cycle not allowed", if you can pass this linter you are ready to go. ENJOY !!!

