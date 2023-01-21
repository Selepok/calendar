FROM golang:1.19.1

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY . .

#RUN go build -o main ./cmd/server/main.go

CMD ["air", "-c", ".air.toml"]
