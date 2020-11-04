FROM alpine:latest

WORKDIR /home/

COPY build/ .
CMD ["./nats-client"]
