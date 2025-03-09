# === Build Stage ===
FROM golang:1.24.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build apontando para o diretório correto
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /go-app ./cmd/server

# === Runtime Stage ===
FROM scratch
COPY --from=builder /go-app /go-app
COPY --from=builder /app/.env /.env
COPY --from=builder /app/infra/database/migrations /infra/database/migrations

ENTRYPOINT ["/go-app"]
