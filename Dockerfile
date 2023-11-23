FROM golang:1.21.3-alpine as go-builder

# Instal tools
RUN apk add --no-cache git

WORKDIR /app

# Copy go mod and sum files.
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the service source code.
COPY . .

# Build
RUN GOOS=linux GOARCH=amd64 go build -o /app/main cmd/server/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=go-builder /app/main .
ENTRYPOINT ["/app/main"]
