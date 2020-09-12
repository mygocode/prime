# Start from the latest golang base image
FROM golang:1.13.8 AS builder

# Add Maintainer Info
LABEL maintainer="Muhammad <inform.tariq@gmail.com.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy Go Modules dependency requirements file
COPY go.mod .

# Copy Go Modules expected hashes file
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy all the app sources (recursively copies files and directories from the host into the image)
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o prime .


## --- Start a new stage from scratch --- ##
FROM alpine:latest  

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app .

# Set http port
EXPOSE 8000

# Build the app
# RUN go build 

# Run the app
CMD ["./prime"]
