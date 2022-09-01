FROM golang:1.19.0-buster as build
MAINTAINER <scorpion@dgu.ac.kr>
# $GOPATH 대신 모듈에 있는 패키지를 사용합니다
ENV GO111MODULE=on

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o tidify-api

EXPOSE 8888

CMD [ "/usr/src/app/tidify-api" ]