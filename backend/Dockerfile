FROM golang:1.10 as builder
WORKDIR /go/src/github.com/Vorian-Atreides/24sessions/backend
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app ./cmd/api

FROM scratch
EXPOSE 8080
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app /app
CMD ["/app"]
