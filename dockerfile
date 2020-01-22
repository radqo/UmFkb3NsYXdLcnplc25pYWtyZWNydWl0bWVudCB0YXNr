FROM golang:1.13.4-alpine3.10

RUN mkdir -p /app/src/github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr

ADD . /app/src/github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr
WORKDIR /app/src/github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr

RUN ls 

ENV GOPATH=/app

RUN apk add --no-cache git mercurial \
    && go get "github.com/gorilla/mux" \
    && apk del git mercurial

RUN go build -o weatherapi /app/src/github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/cmd/server/main.go
CMD ["/app/src/github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/weatherapi"]