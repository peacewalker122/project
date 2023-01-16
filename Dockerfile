FROM golang:1.19-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN apk add curl
RUN  curl -L https://github.com/golang-migrate/migrate/releases/download/v4.12.2/migrate.linux-amd64.tar.gz | tar xvz
RUN go build -o project main.go

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/project .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY db/migration/project ./migration

COPY .env .
COPY wait-for.sh .
COPY start.sh .
COPY migrate.sh .

EXPOSE 8080
CMD ["/app/project"]