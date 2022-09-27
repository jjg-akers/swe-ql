# build stage
FROM golang:1.17-buster as builder

WORKDIR /app

COPY app /app

RUN go mod download

# build binary
RUN go build -o /sweql ./cmd/main.go

# Run stage
FROM debian:buster

RUN apt-get update --fix-missing && \
apt-get install -yqq ca-certificates

COPY --from=builder /sweql .

CMD ["/sweql"]
