FROM golang:1.10

# Install dep
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

COPY ./ /go/src/github.com/mpsido/leboncoin-go/
WORKDIR /go/src/github.com/mpsido/leboncoin-go/
RUN dep ensure -v
ENV PORT 8080

CMD go run main.go