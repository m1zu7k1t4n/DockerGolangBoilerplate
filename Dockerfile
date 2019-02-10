FROM golang:alpine AS build-env
MAINTAINER arasan01
ADD ./src /go/src/server
RUN go install server

FROM alpine
COPY --from=build-env /go/bin/server .
ENV PORT 8080
CMD ["./server"]
