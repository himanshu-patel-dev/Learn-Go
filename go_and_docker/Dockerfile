# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
# download packages listed in go.mod
RUN go mod download
# copy source code
COPY *.go ./
# compile application
RUN go build -o /go-test-app
EXPOSE 8080
# execute the binary
CMD [ "/go-test-app" ]
