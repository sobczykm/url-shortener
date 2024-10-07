# Build stage
FROM golang:1.23-alpine as builder
WORKDIR /build
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main ./main.go

# Final stage
FROM alpine:3
COPY --from=builder /main /bin/main
EXPOSE 3000
ENTRYPOINT ["/bin/main"]