// module rafi-server
module cpfi-server

go 1.21

toolchain go1.21.0

replace cpfi.client => ../client

require cpfi.client v0.0.0

require (
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	golang.org/x/exp v0.0.0-20230321023759-10a507213a29 // indirect
	gonum.org/v1/gonum v0.14.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)
