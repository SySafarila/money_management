FROM golang:1.25.0-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build main.go

FROM alpine:3.20

RUN apk add tzdata

WORKDIR /app

COPY --from=builder /app/main /start

EXPOSE 3000

CMD ["/start"]