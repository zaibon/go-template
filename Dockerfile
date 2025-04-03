# Dockerfile
# Use a multi-stage build to reduce image size and support multi-platform builds

FROM --platform=$BUILDPLATFORM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/service

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app ./

EXPOSE 8080 9090

CMD ["./app"]