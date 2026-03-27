##
# builder
##

FROM golang:1.26.1-alpine3.23 AS builder

WORKDIR /app
COPY . .

RUN go build -o main main.go && chmod +x main

##
# app
##

FROM alpine:3.23.3

WORKDIR /app
COPY --from=builder /app/main /app/

ENTRYPOINT ["/app/main"]
