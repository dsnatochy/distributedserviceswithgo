module github.com/dsnatochy/proglog

go 1.13

require (
	github.com/casbin/casbin v1.9.1
	github.com/cloudflare/cfssl v1.4.1 // indirect
	github.com/golang/protobuf v1.4.1
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/hashicorp/memberlist v0.2.4 // indirect
	github.com/hashicorp/raft v1.1.1
	github.com/hashicorp/raft-boltdb v1.1.1
	github.com/hashicorp/serf v0.8.5
	github.com/soheilhy/cmux v0.1.5
	github.com/stretchr/testify v1.7.0
	github.com/travisjeffery/go-dynaport v1.0.0
	github.com/tysontate/gommap v0.0.0-20210506040252-ef38c88b18e1
	go.opencensus.io v0.22.2
	go.uber.org/zap v1.10.0
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
	google.golang.org/grpc v1.32.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.0.0 // indirect
	google.golang.org/protobuf v1.25.0
)

replace github.com/hashicorp/raft-boltdb => github.com/travisjeffery/raft-boltdb v1.0.0
