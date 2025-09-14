#base image from dockerhub
FROM golang:1.25.0-alpine AS builder

#Set working derictory on container
WORKDIR /app

#Copy file go.mod and go.sum
COPY go.mod go.sum ./
#Download dependency from go.mod and save cache on container
RUN go mod download

#Copy all source code from project on container (/app)
COPY . .

# Build the binary from cmd/main.go
RUN go build -o server ./cmd/main.go

# 2. Create a small final image
FROM alpine:latest
WORKDIR /root/

# Copy binary from builder stage
COPY --from=builder /app/server .

# Expose port (Echo defaults to 1323, unless you changed it)
EXPOSE 1323

# Run the app
CMD ["./server"]