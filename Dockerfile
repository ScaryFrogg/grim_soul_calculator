# Build Stage
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build ./cmd/api/main.go
RUN chmod +x main
# Now create the final image
FROM golang:1.23
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/gs_calc.db .
EXPOSE 3001
CMD ["./main"]
