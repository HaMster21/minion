FROM golang:1.6-alpine

COPY . /go/src/app
WORKDIR /go/src/app

RUN apk add git --update-cache && go get -d -v && go install -v

EXPOSE "5000"
CMD ["app"]
