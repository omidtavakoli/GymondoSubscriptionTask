# README

## Requirements
```

https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
```

## Local PostgreSQL
Use 
```
docker-compose up -d 
```

or 
```
docker run \
  -d \
  -e POSTGRES_HOST_AUTH_METHOD=trust \
  -e POSTGRES_USER=test \
  -e POSTGRES_PASSWORD=123456 \
  -e POSTGRES_DB=gymondo \
  -p 5432:5432 \
  postgres:12.5-alpine
```

## Migrations

Run:

```
migrate -path db/migration -database "postgresql://test:123456@localhost:5432/gymondo?sslmode=disable" up
```

Create:

```
migrate create -ext sql -dir db/migrations/ <migration name>
```
