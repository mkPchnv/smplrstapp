FROM golang:latest as build

ENV GO111MODULE=auto

RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./cmd/smplrstapp/main.go ./
COPY ./.env ./

RUN go build -o ./smplrstapp ./main.go

EXPOSE 8020

CMD [ "./smplrstapp" ]