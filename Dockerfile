FROM golang

COPY ./app /go/src/github.com/user/myProject/app
WORKDIR /go/src/github.com/user/myProject/app

CMD go get ./...
CMD go run server.go

EXPOSE 8080