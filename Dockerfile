FROM golang:1.22

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . ./


# Build the Go app
RUN make build

EXPOSE 8080

CMD ["./utf-moodle-scraper"]