package cpfi

import (
	"log"
	"net/rpc"
)

// const SERVER_ADDR = "172.16.238.1:8080"
// opengauss rqlite
const SERVER_ADDR = "172.17.0.1:8080"

// const SERVER_ADDR = "172.17.0.1:8081"

// const SERVER_ADDR = "127.0.0.1:8080"

var rpcClient *rpc.Client = nil

func initRPCClient() error {
	if rpcClient != nil {
		return nil
	}
	hostPort := SERVER_ADDR
	var err error
	rpcClient, err = rpc.Dial("tcp", hostPort)
	if err != nil {
		log.Printf("error in setting up connection to %s due to %v\n", hostPort, err)
		return err
	}
	return nil
}

func printRPCError(err error) {
	log.Printf("Sieve client RPC error: %v\n", err)
}
