# build stage
FROM iron/go:dev AS build-env
ADD . /src
RUN go get -u github.com/satori/go.uuid
RUN cd /src && go build -o main

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/ /app/
ENTRYPOINT ["./main"]