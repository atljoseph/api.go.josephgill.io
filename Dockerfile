# ############################
# # STEP 1 build executable binary
# ############################
# FROM golang@sha256:cee6f4b901543e8e3f20da3a4f7caac6ea643fd5a46201c3c2387183a332d989 as builder
# # Install git + SSL ca certificates.
# # Git is required for fetching the dependencies.
# # Ca-certificates is required to call HTTPS endpoints.
# RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates
# # Create appuser
# RUN adduser -D -g '' appuser
# WORKDIR $GOPATH/src/mypackage/myapp/
# COPY . .
# # Fetch dependencies.
# # Using go get.
# RUN go get -d -v
# # Using go mod.
# # RUN go mod download
# # RUN go mod verify
# # Build the binary
# RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/hello
# ############################
# # STEP 2 build a small image
# ############################
# FROM scratch
# # Import from builder.
# COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY --from=builder /etc/passwd /etc/passwd
# # Copy our static executable
# COPY --from=builder /go/bin/hello /go/bin/hello
# # Use an unprivileged user.
# USER appuser
# # Run the hello binary.
# ENTRYPOINT ["/go/bin/hello"]

##################################################


# # Dockerfile References: https://docs.docker.com/engine/reference/builder/

# # Start from the latest golang base image
# FROM golang:alpine as builder

# # Add Maintainer Info
# LABEL maintainer="Joseph Gill <joseph.gill.atlanta@gmail.com>"

# # install extras
# RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# # add unpriveledged user
# RUN adduser -D -g '' appuser

# # Set the Current Working Directory inside the container
# WORKDIR /app

# # Copy go mod and sum files
# COPY go.mod go.sum ./

# # Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
# RUN go mod download

# # Copy the source from the current directory to the Working Directory inside the container
# COPY . .

# # Build the Go app
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# ######## Start a new stage from scratch #######
# FROM alpine:latest as final

# WORKDIR /root/

# # Copy the Pre-built binary file from the previous stage
# # also other important info
# COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY --from=builder /etc/passwd /etc/passwd
# COPY --from=builder /app/main .

# # # Use an unprivileged user.
# USER appuser

# # Expose port 8080 to the outside world
# EXPOSE 8080

# # Command to run the executable
# CMD ["./main"] 

#################################################

# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:alpine as builder

# Add Maintainer Info
LABEL maintainer="Joseph Gill <joseph.gill.atlanta@gmail.com>"

# RUN apk --no-cache add ca-certificates tzdata openssl
# # && update-ca-certificates

# RUN openssl verify /etc/ssl/certs/ca-certificates.crt
# RUN openssl ecparam -genkey -name secp384r1 -out /etc/ssl/private/server.key
# RUN openssl req -new -x509 -sha256 -key /etc/ssl/private/server.key -out /etc/ssl/certs/server.crt -days 3650


# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
# COPY https-server.key https-server.crt ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


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