FROM golang:1.22-alpine AS build
WORKDIR /app

# Copia apenas o go.mod; go.sum pode não existir inicialmente
COPY go.mod ./
RUN go mod download

COPY . .

RUN go mod tidy

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o http ./cmd/api

FROM alpine:3.20
WORKDIR /app

COPY --from=build /app/http .

EXPOSE 8080

CMD ["./http"]