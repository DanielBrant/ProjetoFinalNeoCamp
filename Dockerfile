FROM golang:1.22-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o http .

FROM alpine:3.20
WORKDIR /app

COPY --from=build /app/http .

EXPOSE 8080

CMD ["./http"]