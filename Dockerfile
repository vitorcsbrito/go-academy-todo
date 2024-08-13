FROM golang:1.22.5-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

ADD ./ .

RUN go build -o /go-docker-demo

EXPOSE 8080

CMD [ "/go-docker-demo" ]