FROM golang:1.15 AS builder

WORKDIR /source
COPY ./api/go.mod .
COPY ./api/go.sum .
# Download dependencies
RUN go mod download

COPY ./api/ .
# Compile api
RUN CGO_ENABLED=0 GOOS=linux go build -a -o airquality .

# Remove golang dependencies
FROM alpine:latest
WORKDIR /airquality/
COPY --from=builder /source/airquality .
# Run api
CMD ["./airquality"]