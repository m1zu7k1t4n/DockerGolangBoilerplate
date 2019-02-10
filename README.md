# Docker Golang Boilerplate

## Dockerfile

Selected Method of "Multi Stage Building" by Dockerfile.

'build-env stage' compile stage.
Compile the project in the src directory.
It contains the results built into the last container.


example :
```Dockerfile
FROM golang:alpine AS build-env
MAINTAINER arasan01
ADD ./src /go/src/server
RUN go install server

FROM alpine
COPY --from=build-env /go/bin/server .
ENV PORT 8080
CMD ["./server"]
```

## Script

build and run in the script directory.
