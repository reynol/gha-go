FROM golang:1.21

ENV PATH_MAIN=""

COPY . /home/docker/go/src

WORKDIR /home/docker/go/src$PATH_MAIN

RUN go build -v -o /home/docker/go/build .

FROM alpine:3.17
COPY --from=0 /home/docker/go/build ./go/build

CMD ["./go/build"]