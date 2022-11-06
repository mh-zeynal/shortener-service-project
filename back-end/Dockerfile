FROM golang:1.16-alpine
ADD . /go/src/back-end
WORKDIR /go/src/back-end
RUN go get back-end
RUN go install
ENTRYPOINT ["/go/bin/back-end"]
