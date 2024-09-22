FROM golang:1.22-alpine3.20 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -ldflags="-s -w" -o main main.go

# run stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/etc ./etc
COPY start.sh .
RUN chmod +x /app/start.sh
EXPOSE 8888

ENTRYPOINT [ "/app/start.sh" ]