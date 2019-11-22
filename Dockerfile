
# TODO: Dockerfile for prod

# Start from the latest golang base image
FROM golang:1.13-alpine as builder

# Add Maintainer Info
LABEL maintainer="Joseph Gill <joseph.gill.atlanta@gmail.com>"

# update
RUN apk --update upgrade

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

######## Start a new stage from scratch #######
FROM alpine:latest as final

# DEV ONLY #
# add bash et al
RUN apk add --no-cache bash coreutils grep sed curl

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Create an unprivileged user (order of operations matters)
RUN adduser --disabled-password --gecos '' appuser

# Do not forget to chmod so no permission errors
RUN chmod 755 .

# Use the unpriveleged user
USER appuser

# Expose port 8080 to the outside world
# EXPOSE 8080

# Command to run the executable
# CMD ["./main"] 
ENTRYPOINT ./main