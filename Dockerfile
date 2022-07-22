FROM golang:1.17-alpine3.16 as builder

RUN mkdir /app
WORKDIR /app
COPY . .

RUN go build -o alta-mnroom

FROM alpine
WORKDIR /app
COPY --from=builder /app/ /app/
CMD ["./alta-mnroom"]