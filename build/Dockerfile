# Certificates
FROM alpine:latest as certs
RUN apk --update add ca-certificates

# Build
FROM golang:1.15 as builder

ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR /code
COPY . .
RUN go build -o /controller cmd/manager/main.go

# Final image
FROM scratch

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /controller /controller

ENTRYPOINT ["/controller"]
