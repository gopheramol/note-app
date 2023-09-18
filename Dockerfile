FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0 
RUN go build -o main .

FROM alpine:latest
COPY --from=builder /app ./

EXPOSE 8081
ENTRYPOINT ["./main"]
