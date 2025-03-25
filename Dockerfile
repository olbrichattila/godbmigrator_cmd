FROM golang:1.22-alpine AS builder
WORKDIR /app

ENV CGO_ENABLED=1

RUN apk add --no-cache git \
    && apk add --no-cache gcc musl-dev sqlite-dev

COPY . .
RUN go mod tidy
RUN go build ./cmd/migrator

# Final image
FROM alpine:latest
WORKDIR /app

# Install SQLite runtime dependencies
RUN apk add --no-cache sqlite-libs

COPY --from=builder /app/migrator /app/migrator

EXPOSE 8080
CMD ["/app/migrator", "serve"]
