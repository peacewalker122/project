FROM golang:1.19-alpine3.17 AS builder
WORKDIR /app
COPY . . 
RUN go build -o project main.go

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/project .

EXPOSE 8080
CMD ["./app"]