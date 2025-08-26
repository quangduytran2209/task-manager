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

#Buiding Go source to binary main
RUN go build -o server

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .
CMD ["./server"]