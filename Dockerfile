# build stage
FROM golang:alpine AS build-env
WORKDIR /go/src/stage
ADD . /go/src/stage
RUN apk add --no-cache git mercurial && go get ./...
RUN go build -o main

# ENTRYPOINT ["./main"]

# run stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/stage/main /app/
ENTRYPOINT ./main