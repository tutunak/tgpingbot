FROM golang:alpine3.18 as builder
LABEL authors="tutunak"

COPY . /app
WORKDIR /app
RUN go build -o goapp .

FROM alpine:3.18 as production
LABEL authors="tutunak"
COPY --from=builder /app/goapp /app/goapp
RUN addgroup -S feedcheck && adduser -S feedcheck -G feedcheck && \
    chown -R goapp:goapp /app
USER goapp
WORKDIR /app
CMD ["./feedcheck"]
