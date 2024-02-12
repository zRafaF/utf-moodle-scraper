# Stage 1: Build Stage
FROM golang:1.22 AS build

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

# Stage 2: Final Stage
FROM scratch

# Copy the binary from the Build Stage into the Final Stage
COPY --from=build /app/utf-moodle-scraper /utf-moodle-scraper

EXPOSE 8080

# Command to run the executable
CMD ["/utf-moodle-scraper"]
