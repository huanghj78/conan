module reproduce

go 1.19

replace cpfi.client => ../client

require cpfi.client v0.0.0-00010101000000-000000000000

require (
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	go.etcd.io/raft/v3 v3.0.0-20231012085229-7c3ed830bbb0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
)
