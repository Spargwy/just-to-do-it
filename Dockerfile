FROM golang:1.21-alpine as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ./tasker ./app/client/cmd/
EXPOSE 3000

CMD ["./tasker"]

