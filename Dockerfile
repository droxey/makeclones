
# Base build image
FROM golang:1.11-alpine

# Install dependencies needed to build the project
RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /go/src/
RUN mkdir makeclones
WORKDIR /go/src/makeclones

# Force the go compiler to use modules
ENV GO111MODULE=on

# Populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

# This is the ‘magic’ step that will download all the dependencies that are specified in
# the go.mod and go.sum file.
# Because of how the layer caching system works in Docker, the go mod download
# command will _ only_ be re-run when the go.mod or go.sum file change
# (or when we add another docker instruction this line)
RUN go mod download

# Copy the rest of the source code
COPY . .

# And compile the project!
ENV CGO_ENABLED=1
RUN GOOS=linux GOARCH=amd64 go build github.com/droxey/makeclones
RUN GOOS=darwin GOARCH=amd64 go build github.com/droxey/makeclones
