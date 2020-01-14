FROM golang:1.13.4-alpine3.10
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o weatherapi .
CMD ["/app/weatherapi"]