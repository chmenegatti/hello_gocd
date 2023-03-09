# Build stage
FROM golang:latest AS build

# Set workdir
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy all other source files
COPY . .

# Compile binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

# Run stage
FROM alpine

# Set workdir and copy binary
WORKDIR /app
COPY --from=build /app/server .
COPY --from=build /.env .

# Expose port 5000
EXPOSE 5000

# Start server
CMD ["./server"]
