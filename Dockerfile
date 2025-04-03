# Dockerfile
# Use a multi-stage build to reduce image size and support multi-platform builds

FROM --platform=$BUILDPLATFORM golang:1.24-alpine AS builder
ARG TARGETOS
ARG TARGETARCH
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o app ./cmd/service

FROM alpine:latest
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
WORKDIR /app

COPY --from=builder /app/app ./

# Switch to a non-root user for security
USER appuser

EXPOSE 8080 9090

CMD ["./app"]