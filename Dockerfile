FROM golang:alpine as build

ENV GOPATH=/

WORKDIR /app

COPY pkg pkg
COPY src src
COPY go.mod .
COPY go.sum .

RUN go get -d ./...
RUN go build src/main.go

FROM alpine:latest
COPY --from=build ./app/main .
CMD [ "./main" ]