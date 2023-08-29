FROM golang:1.21

ENV PATH_MAIN="/cmd/fraud"

COPY . /home/docker/go/src

WORKDIR /home/docker/go/src$PATH_MAIN

RUN go build -v -o /home/docker/go/build .

CMD ["./go/build"]