FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o auth_service .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/auth_service /auth_service

CMD ["/auth_service"]

EXPOSE 8080