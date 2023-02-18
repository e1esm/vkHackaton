FROM golang:1.17-alpine
WORKDIR /app

RUN apk update && apk add libc-dev && apk add gcc && apk add make && apk add git && apk add bash

COPY ./go.mod go.sum ./
RUN go mod download && go mod verify

ENV GOBIN /go/bin

RUN go get github.com/githubnemo/CompileDaemon


COPY ./* /app/
COPY ./entrypoint.sh /entrypoint.sh
COPY ./.env /.env


ADD https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /usr/local/bin/wait-for
RUN chmod +rx /usr/local/bin/wait-for /entrypoint.sh

ENTRYPOINT ["sh", "/entrypoint.sh"]