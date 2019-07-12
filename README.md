# gomplex

**Performs math operations on complex numbers**

## gRPC service
### Install protoc
`$ go get -u github.com/golang/protobuf/protoc-gen-go`

### Generate gRPC code in GO
`protoc -I service service/complex.proto --go_out=plugins=grpc:service`

### Build & run
`$ go build app/`

`$ ./app`

### Docker
Via docker cli
`$ docker run gomplex`

Via docker compose `$ docker-compose up`


### Kubernetes

## Clients

### Go client
#### Checking complex number type
`$ go run clients/client.go`
```
Real part: 2
Imaginary part: -1
Type of 
{
        Real:    2.00 
        Imaginary:       -1.00
}
is Complex
```

### Python client

#### Installation
##### Install requirements
`$ pip3 install -r requirements.txt`

##### Generate gRPC code
`$ python3 -m grpc_tools.protoc -Iservice --python_out=clients/ --grpc_python_out=clients/ service/service.proto`

#### Running example
`$ python3 clients/client.py --number 1+j --number 2-3j`
`$ python3 clients/client.py  --number 2-j --number=-1-4j --subtraction`


## TO-DO
