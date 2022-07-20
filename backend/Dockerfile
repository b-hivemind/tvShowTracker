FROM golang:1.16-alpine

ENV TZ=America/Los_Angeles
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENV GIN_MODE=release

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

ADD pkg/ ./pkg

RUN go build -o preparer

EXPOSE 10090
CMD ["/app/preparer", "-mode", "live"]
