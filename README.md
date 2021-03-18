This repo used for reproduce panic issues.

Just run `go run *.go`

Tools
```
need install protoc tools from https://github.com/protocolbuffers/protobuf/releases
```

```
GO111MODULE=on go get -v -u \
    github.com/mitchellh/protoc-gen-go-json \
    google.golang.org/protobuf/cmd/protoc-gen-go    
```
