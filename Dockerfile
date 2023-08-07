FROM golang:alpine as build

ENV GOPATH=/

WORKDIR /app

COPY pkg pkg
COPY cmd cmd
COPY internal internal
COPY go.mod .
COPY go.sum .

RUN go get -d ./...
RUN go build cmd/main.go

FROM alpine:latest
COPY --from=build ./app/main .
CMD [ "./main" ]