
# Start from the latest golang base image
FROM golang:alpine as builder

# Add Maintainer Info
LABEL maintainer="Joseph Gill <joseph.gill.atlanta@gmail.com>"

# update
RUN apk --update upgrade

# enable usage of cgo
# See http://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
# RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN apk add gcc g++ 
#--no-cache 

# add sqlite
RUN apk add sqlite

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
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main .


######## Start a new stage from scratch #######
FROM alpine:latest as final

WORKDIR /root/

# RUN apk --no-cache add ca-certificates tzdata 
# # && update-ca-certificates

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy other important files
# COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
# COPY --from=builder /etc/ssl/certs/server.crt /etc/ssl/certs/
# COPY --from=builder /etc/ssl/private/server.key /etc/ssl/private/
# COPY --from=builder /app/https-server.key .
# COPY --from=builder /app/https-server.crt .
# https://github.com/denji/golang-tls

# Create an unprivileged user
# Do not forget to chmod so no permission errors
RUN adduser --disabled-password --gecos '' appuser
RUN chmod 755 .

# Use the unpriveleged user
USER appuser

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"] 