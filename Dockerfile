# build go in docker
FROM golang:1.9.1-alpine3.6 AS build-env
ADD . /go/src/app
RUN cd /go/src/app && go build -o goapp

# build an image with the compiled binary
FROM alpine:3.6
RUN apk update && apk add ca-certificates
WORKDIR /app
COPY --from=build-env /go/src/app/goapp /app/
ENTRYPOINT ./goapp