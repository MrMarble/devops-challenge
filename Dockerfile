FROM golang:1.14 AS builder

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

# Needed for "pg_isready" in order to determine if the database is ready to accept connections
RUN apk --no-cache add postgresql-client

WORKDIR /airquality/

COPY --from=builder /source/airquality .

# Run api
CMD ["./airquality"]