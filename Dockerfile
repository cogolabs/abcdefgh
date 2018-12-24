FROM golang:1.11
ADD . /go/src/github.com/cogolabs/abcdefgh
RUN go get -x github.com/cogolabs/abcdefgh
CMD ["abcdefgh", "--help"]