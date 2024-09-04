FROM golang:1.23-alpine

WORKDIR /usr/src/app

RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xvz

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -v -o /usr/local/bin/app ./cmd/api
RUN chmod +x ./build/docker/go/migrate.sh

ENTRYPOINT ["./build/docker/go/migrate.sh"]

CMD ["app"]