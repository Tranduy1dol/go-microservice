FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /api ./cmd/api
RUN CGO_ENABLED=0 go build -o /importer ./cmd/importer

FROM alpine:3.21
RUN apk add --no-cache ca-certificates tzdata
COPY --from=builder /api /usr/local/bin/api
COPY --from=builder /importer /usr/local/bin/importer
COPY config/config.yaml /etc/learning-japanese/config.yaml
EXPOSE 8080
CMD ["api"]
