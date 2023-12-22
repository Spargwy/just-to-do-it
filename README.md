# just-to-do-it
simple task manager

## Prerequests: 
- openssl(optional)
- Go language(optional)
- docker/docker-compose


```git clone git@github.com:Spargwy/just-to-do-it.git```

For authorization logic, program need pair of rsa keys.
You can use existing app.rsa and app.rsa.pub that are locates in root of the project. 
But for security in production you should generate them with openssl(may need to be installed):

```openssl genrsa -out /opt/keys/tasker.rsa -outform PEM -pubout -out tasker.rsa.pub```

```openssl rsa -in /opt/keys/tasker.rsa -outform PEM -pubout -out /opt/keys/tasker.rsa.pub```

docker and project configuration locates in .env file. Configuration options can be find in `config.go` file You can use content from .env.example for start:

```cat .env.example > .env```


start db with 

```cd docker && docker-compose up -d ```


apply migrations
```cat db/schema.sql | psql postgres://tasker:password@localhost:5432/todo```


Run project with

``` go run app/client/cmd/main.go```

Or via docker:
```docker-compose up -d```