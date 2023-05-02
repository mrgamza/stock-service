FROM golang:1.15
MAINTAINER outofcoding@gmail.com

ENV GOPATH /go
ARG service=go-service

WORKDIR $GOPATH/src/${service}

COPY go.mod .
COPY go.sum .
RUN go mod download

ADD . /go/src/${service}

EXPOSE 5000

CMD ["go", "run", "go-service"]