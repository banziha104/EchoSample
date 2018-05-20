# Docker로 배포

- main 실행 파일 만듬 : CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main . 
- dockerfile 작성 

```dockerfile
#!/bin/bash

FROM golang

MAINTAINER Devign92 <banziha104@gmail.com>


ADD main /
CMD ["/main"]
EXPOSE 80
```

- docker 빌드 : docker build -t dldudwnsdl/sample:0.1 . <name>:<tag> <path>
- dockerhub에 푸쉬 : docker push dldudwnsdl/sample:0.1
- 서버에서 ec2 도커 이미지 설치 및 구동 : docker run -p 80:80 dldudwnsdl/sample:0.1  