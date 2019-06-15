FROM golang:1.10

# Install dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

COPY ./ /go/src/github.com/mpsido/leboncoin-go/
WORKDIR /go/src/github.com/mpsido/leboncoin-go/
RUN dep ensure -v

CMD go run main.go