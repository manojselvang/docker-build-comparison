FROM golang:1.22

WORKDIR /app

COPY app/go.mod ./

RUN go mod download

COPY app/ .

RUN go build -o server .

EXPOSE 8080

CMD ["./server"]

