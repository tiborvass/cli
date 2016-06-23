FROM golang:1.6.2
RUN mkdir -p /go/src/github.com/cpuguy83
RUN mkdir -p /go/src/github.com/cpuguy83 \
    && git clone -b v1.0.5 https://github.com/cpuguy83/go-md2man.git /go/src/github.com/cpuguy83/go-md2man \
    && cd /go/src/github.com/cpuguy83/go-md2man \
    && go get -v ./...
CMD ["/go/bin/go-md2man", "--help"]
