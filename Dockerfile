FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Create appuser.
RUN adduser -D -g '' appuser

WORKDIR $GOPATH/src/sorting-word
COPY . .

# Using go get.
RUN go get

# building apps in sorting-word
RUN go build -o sorting-word

# running sorting-word
ENTRYPOINT ./sorting-word

# running in port
EXPOSE 8002