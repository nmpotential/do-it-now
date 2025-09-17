# filepath: /Users/nmpotential/workspaces/go/src/github.com/nmpotential/do-it-now/backend/Dockerfile
# Build stage
FROM golang:1.22.2-alpine3.19 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o doitnow -ldflags="-s -w" ./cmd

# Run stage
FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/doitnow .
COPY migrations ./migrations

EXPOSE 8080

CMD ["./doitnow"]