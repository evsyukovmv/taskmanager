# The back-end (REST API) part of the task management application

## Demo app
[https://go-taskmanager.herokuapp.com](https://go-taskmanager.herokuapp.com)

## API Documentation
[https://gotaskmanager.docs.apiary.io/](https://gotaskmanager.docs.apiary.io/)

## Code Status
[![CircleCI](https://circleci.com/gh/evsyukovmv/taskmanager.svg?style=svg)](https://circleci.com/gh/evsyukovmv/taskmanager)

## How to use

### Using Docker & Docker Compose

```
docker-compose up
```

### Without Docker

* Create PostgresSQL database
```sql
CREATE DATABASE taskmanager;
```

* Add environment variables
```sh
export PORT=8080
export DATABASE_URL=postgres://user:password@localhost/taskmanager?sslmode=disable
```

* Run migrations using [migrate](https://github.com/golang-migrate/migrate)
```sh
migrate -path db/migrations -database $DB_URL up
```

* Run application
```sh
go run main.go
```
