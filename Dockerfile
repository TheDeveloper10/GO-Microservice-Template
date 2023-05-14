FROM golang:1.20.4-bullseye as builder

WORKDIR /app

COPY . /app

RUN make get_dependencies && make build && chmod +x bin/main



FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/main /app

CMD [ "/app/main" ]
