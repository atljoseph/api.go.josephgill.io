
# TODO: Dockerfile for prod

# Start from the latest golang base image
FROM golang:1.13-alpine as builder

# Add Maintainer Info
LABEL maintainer="Joseph Gill <joseph.gill.atlanta@gmail.com>"

# update
RUN apk --update upgrade

# enable usage of cgo
# See http://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
# RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
# ENV PATH "/lib:$PATH"
# RUN apk add --no-cache gcc g++ 
RUN apk add --no-cache gcc musl-dev

# add sqlite
RUN apk add --no-cache sqlite

# add bash et al
RUN apk add --no-cache bash coreutils grep sed

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

# Copy in the database
# TODO handle the database turnover better... can db be written to?
COPY photos.db ./

# Create an unprivileged user (order of operations matters)
RUN adduser --disabled-password --gecos '' appuser

# Do not forget to chmod so no permission errors
RUN chmod 755 .

# Use the unpriveleged user
USER appuser

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"] 