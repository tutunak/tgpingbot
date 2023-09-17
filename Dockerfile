FROM golang:alpine3.18 as builder
LABEL authors="tutunak"

COPY . /app
WORKDIR /app
RUN go build -o tgpingbot .

FROM alpine:3.18 as production
LABEL authors="tutunak"
COPY --from=builder /app/tgpingbot /app/tgpingbot
RUN addgroup -S tgpingbot && adduser -S tgpingbot -G tgpingbot && \
    chown -R tgpingbot:tgpingbot /app
USER tgpingbot
WORKDIR /app
CMD ["./tgpingbot"]
