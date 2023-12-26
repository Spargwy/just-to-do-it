# just-to-do-it
simple task manager

## Prerequests: 
- openssl(optional)
- Go language(optional)
- docker/docker-compose
- Postgres

```git clone git@github.com:Spargwy/just-to-do-it.git``` if u using ssh authentication and 
```https://github.com/Spargwy/just-to-do-it``` 
if not

For authorization logic, program need pair of rsa keys.
You can use existing app.rsa and app.rsa.pub that are locates in root of the project. 
But for security in production you should generate them with openssl(may need to be installed):

```openssl genrsa -out ./app.rsa```

```openssl rsa -in ./app.rsa -outform PEM -pubout -out ./app.rsa.pub```

docker and project configuration locates in .env file. Configuration options can be found in `config.go` file. You can use content from .env.example for start:

```cat .env.example > .env```

start tasker and db
```docker-compose up -d```


apply migrations
```cat db/schema.sql | psql postgres://tasker:password@localhost:5432/todo```


Also, you can run it manually if u have go

``` go run app/client/cmd/main.go```


Testing:
- start test db: 
```docker-compose -f postgres.test.yml up -d```
- apply migrations: 
```cat db/schema.sql | psql postgres://tasker:password@localhost:7232/todo```

```go test ./...```