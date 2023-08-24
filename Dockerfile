FROM golang:1.16

WORKDIR /app

COPY . .

RUN go mod download && go mod verify

RUN go build -o main .

CMD ["./main"]