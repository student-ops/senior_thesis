FROM golang:1.18-alpine3.16 AS builder

RUN mkdir /app
WORKDIR /app

COPY ./cmd .
COPY ./internal .

RUN go mod tidy
RUN go build -o main ./

FROM alpine:3.16

COPY --from=builder /app/main .
CMD ["/main"]