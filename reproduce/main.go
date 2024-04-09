package main

import (
	"os"

	client "cpfi.client"
)

// 1710308823
func main() {
	args := os.Args
	path := args[1]
	client.StartWorking(path)
}
