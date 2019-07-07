
### Install protoc
`$ go get -u github.com/golang/protobuf/protoc-gen-go`


### Generate gRPC code in GO
`protoc -I service service/complex.proto --go_out=plugins=grpc:service`
