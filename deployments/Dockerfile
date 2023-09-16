FROM golang as builder

WORKDIR /app

COPY .. .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o go-challenge .

FROM alpine

WORKDIR /app

COPY --from=builder /app/go-challenge .

EXPOSE 3000

CMD ["./go-challenge", "start"]