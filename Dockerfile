FROM golang:rc-alpine

WORKDIR src/gomplex
ENV GOPATH /go/
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin

RUN apk add protobuf git

RUN go get -u github.com/golang/protobuf/protoc-gen-go
RUN go get -u gopkg.in/yaml.v2
RUN go get -u google.golang.org/grpc


COPY . .

RUN protoc -I service service/service.proto --go_out=plugins=grpc:service

RUN go build -i -o gomplex ./app/

CMD ./gomplex
