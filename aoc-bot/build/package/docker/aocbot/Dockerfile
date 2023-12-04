# STEP 1 - Build the binary

# use the golang image as build image
FROM golang:1.21.4 AS builder

# build arguments in order to reuse it for arm containers
ARG GOOS=linux
ARG GOARCH=amd64
ARG GOARM

# activate go modules
ENV GO111MODULE on

# copy the local package files to the container's workspace.
COPY . /go/src/github.com/peek1e/aoc-leadboard-bot

# set the working directory to build the application
WORKDIR /go/src/github.com/peek1e/aoc-leadboard-bot

# compile the program
RUN CGO_ENABLED=0 GOOS="${GOOS}" GOARCH="${GOARCH}" go build \
    -ldflags="-s -w" \
    -a \
    -installsuffix cgo \
    -o /go/bin/aoc-leadboard-bot \
    /go/src/github.com/peek1e/aoc-leadboard-bot/cmd/aocbot/main.go

# STEP 2 - get certificates

FROM alpine:3.18 AS certs

RUN apk add --no-cache \
        ca-certificates=20230506-r0

# STEP 3 - Build a minimal container

# start from scratch
FROM scratch

# copy the static executable
COPY --from=builder /go/bin/aoc-leadboard-bot /aoc-leadboard-bot
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# define the entrypoint
ENTRYPOINT ["/aoc-leadboard-bot"]
