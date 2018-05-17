FROM golang:onbuild


COPY ./app /go/src/github.com/user/myProject/app
WORKDIR /go/src/github.com/user/myProject/app


RUN go get ./
RUN go build

EXPOSE 8080