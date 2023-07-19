FROM golang:alpine

ENV GOPATH=/

WORKDIR /app

COPY pkg pkg
COPY src src
COPY go.mod .
COPY go.sum .

RUN go get -d ./...
RUN go build src/main.go

CMD [ "./main" ]