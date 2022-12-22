# Builder
FROM golang:1.19 as builder

# WORKDIR name should be initia-apis
WORKDIR /initia-apis

COPY . .
RUN go build -o build/initia-apis main.go

# Runner
FROM debian:bullseye

WORKDIR /initia-apis

COPY --from=builder /initia-apis/build/initia-apis ./build/initia-apis
COPY --from=builder /initia-apis/.env .

EXPOSE 8999

CMD ["./build/initia-apis"]