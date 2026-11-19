FROM golang:1.22-alpine as builder

RUN mkdir /app
WORKDIR /app
COPY . .

RUN go build -o alta-mnroom

FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/ /app/
CMD ["./alta-mnroom"]