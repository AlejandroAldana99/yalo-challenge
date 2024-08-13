# Dockerfile
FROM golang:1.21-alpine

WORKDIR /usr/src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod tidy

COPY . .

EXPOSE 8000

CMD [ "go run ./main.go" ]
