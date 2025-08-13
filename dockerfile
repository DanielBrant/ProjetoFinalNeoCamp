FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o http .

EXPOSE 8080

CMD ["./http"]