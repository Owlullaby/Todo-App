FROM golang:1.16-alpine as builder

WORKDIR /app

ADD . /app

RUN go build -o /todoapp

EXPOSE 4000

CMD [ "/todoapp" ]