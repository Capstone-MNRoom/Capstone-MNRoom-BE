FROM golang:1.17-alpine3.16

RUN mkdir /app

WORKDIR /app

COPY . .

RUN go build -o alta-mnroom

CMD ["./alta-mnroom"]
