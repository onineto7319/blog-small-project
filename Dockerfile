FROM golang:1.15

WORKDIR /blog-small-project
ADD . /blog-small-project
ENV GO111MODULE=on

RUN cd /blog-small-project
RUN go mod download
RUN go build

EXPOSE 8080

ENTRYPOINT ./blog-small-project