FROM golang:alpine AS builder
WORKDIR /app
COPY . ./
RUN go build -o app cmd/api_server/main.go
FROM alpine
WORKDIR /app
COPY --from=builder /app/app ./
COPY --from=builder /app/.env ./
COPY --from=builder /app/config /app/config
COPY --from=builder /app/api /app/api
EXPOSE 6969
CMD [ "./app" ]