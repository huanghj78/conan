package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	dialTimeout    = 200 * time.Second
	requestTimeout = 10 * time.Second
)

func main() {
	// ctx, _ := context.WithTimeout(context.Background(), requestTimeout)
	//客户端配置
	config := clientv3.Config{
		Endpoints:   []string{"172.16.238.100:2379", "172.16.238.101:2379", "172.16.238.102:2379"},
		DialTimeout: dialTimeout,
	}

	//建立连接
	client, err := clientv3.New(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("connect success")
	defer client.Close()

	resp, err := client.Txn(context.Background()).
		If(clientv3.Compare(clientv3.Version("1").WithPrefix(), ">", 0)).
		Then(
			// clientv3.OpPut("1", "1"),
			clientv3.OpGet("1", clientv3.WithRange("3")),
		).
		Else(clientv3.OpGet("1", clientv3.WithRange("3"))).
		Commit()
	if err != nil {
		fmt.Println(err)
	}
	// Check if the transaction was successful
	if resp != nil {
		fmt.Println(resp)
		// Process the results of the Get operation
		// for _, ev := range resp.Responses[0].GetResponseRange().Kvs {
		// 	fmt.Printf("Key: %s, Value: %s\n", ev.Key, ev.Value)
		// }
	} else {
		fmt.Println("Transaction failed.")
	}
	// args := os.Args
	// mode := "Put"
	// if len(args) == 2 {
	// 	mode = "Get"
	// }
	// timeoutCtx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()
	// if mode == "Put" {
	// 	// 创建 etcd 客户端连接
	// 	config := clientv3.Config{
	// 		Endpoints:   []string{"http://172.16.238.100:2379", "http://172.16.238.101:2379", "http://172.16.238.102:2379"}, // etcd 服务器地址
	// 		DialTimeout: 5 * time.Second,
	// 	}

	// 	client, err := clientv3.New(config)
	// 	if err != nil {
	// 		fmt.Println("无法连接到 etcd 服务器:", err)
	// 		os.Exit(1)
	// 	}
	// 	defer client.Close()

	// 	_, err = client.Status(timeoutCtx, config.Endpoints[0])
	// 	if err != nil {
	// 		fmt.Println("Connect timeout")
	// 		os.Exit(1)
	// 	}

	// 	for i := 0; i < 10; i++ {
	// 		key := fmt.Sprintf("k%d", i)
	// 		value := fmt.Sprintf("v%d", i)
	// 		_, err = client.Put(context.Background(), key, value)
	// 		if err != nil {
	// 			fmt.Println("无法存储值:", err)
	// 			os.Exit(1)
	// 		}
	// 		_, err = client.Status(timeoutCtx, config.Endpoints[0])
	// 		if err != nil {
	// 			fmt.Println("Connect timeout")
	// 			os.Exit(1)
	// 		}
	// 	}

	// 	os.Exit(0)
	// } else {
	// 	config1 := clientv3.Config{
	// 		Endpoints:   []string{"http://172.16.238.100:2379"}, // etcd 服务器地址
	// 		DialTimeout: 5 * time.Second,
	// 	}
	// 	client1, err := clientv3.New(config1)
	// 	if err != nil {
	// 		fmt.Println("无法连接到node1:", err)
	// 		os.Exit(1)
	// 	}
	// 	_, err = client1.Status(timeoutCtx, config1.Endpoints[0])
	// 	if err != nil {
	// 		fmt.Println("Connect timeout")
	// 		os.Exit(1)
	// 	}
	// 	config2 := clientv3.Config{
	// 		Endpoints:   []string{"http://172.16.238.101:2379"}, // etcd 服务器地址
	// 		DialTimeout: 5 * time.Second,
	// 	}
	// 	client2, err := clientv3.New(config2)
	// 	if err != nil {
	// 		fmt.Println("无法连接到node2:", err)
	// 		os.Exit(1)
	// 	}
	// 	_, err = client2.Status(timeoutCtx, config2.Endpoints[0])
	// 	if err != nil {
	// 		fmt.Println("Connect timeout")
	// 		os.Exit(1)
	// 	}
	// 	config3 := clientv3.Config{
	// 		Endpoints:   []string{"http://172.16.238.102:2379"}, // etcd 服务器地址
	// 		DialTimeout: 5 * time.Second,
	// 	}
	// 	client3, err := clientv3.New(config3)
	// 	if err != nil {
	// 		fmt.Println("无法连接到node3:", err)
	// 		os.Exit(1)
	// 	}
	// 	_, err = client3.Status(timeoutCtx, config3.Endpoints[0])
	// 	if err != nil {
	// 		fmt.Println("Connect timeout")
	// 		os.Exit(1)
	// 	}
	// 	for n := 0; n < 10; n++ {
	// 		key := fmt.Sprintf("k%d", n)
	// 		value := fmt.Sprintf("v%d", n)
	// 		resp1, err := client1.Get(context.Background(), key)
	// 		if err != nil {
	// 			fmt.Println("无法获取值:", err)
	// 			os.Exit(1)
	// 		}
	// 		resp2, err := client2.Get(context.Background(), key)
	// 		if err != nil {
	// 			fmt.Println("无法获取值:", err)
	// 			os.Exit(1)
	// 		}
	// 		resp3, err := client3.Get(context.Background(), key)
	// 		if err != nil {
	// 			fmt.Println("无法获取值:", err)
	// 			os.Exit(1)
	// 		}
	// 		_, err = client1.Status(timeoutCtx, config1.Endpoints[0])
	// 		if err != nil {
	// 			fmt.Println("Connect timeout")
	// 			os.Exit(1)
	// 		}
	// 		_, err = client2.Status(timeoutCtx, config2.Endpoints[0])
	// 		if err != nil {
	// 			fmt.Println("Connect timeout")
	// 			os.Exit(1)
	// 		}
	// 		_, err = client3.Status(timeoutCtx, config3.Endpoints[0])
	// 		if err != nil {
	// 			fmt.Println("Connect timeout")
	// 			os.Exit(1)
	// 		}
	// 		for _, kv := range resp1.Kvs {
	// 			fmt.Println(string(kv.Value))
	// 			if value != string(kv.Value) {
	// 				fmt.Println("err value", string(kv.Value))
	// 				os.Exit(1)
	// 			}
	// 		}
	// 		for _, kv := range resp2.Kvs {
	// 			fmt.Println(string(kv.Value))
	// 			if value != string(kv.Value) {
	// 				fmt.Println("err value", string(kv.Value))
	// 				os.Exit(1)
	// 			}
	// 		}
	// 		for _, kv := range resp3.Kvs {
	// 			fmt.Println(string(kv.Value))
	// 			if value != string(kv.Value) {
	// 				fmt.Println("err value", string(kv.Value))
	// 				os.Exit(1)
	// 			}
	// 		}
	// 	}
	// }
	// os.Exit(0)

	// 存储值
	// _, err = client.Put(context.Background(), key, value)
	// if err != nil {
	// 	fmt.Println("无法存储值:", err)
	// 	return
	// }

	// 获取值
	// resp, err := client.Get(context.Background(), key)
	// if err != nil {
	// 	fmt.Println("无法获取值:", err)
	// 	return
	// }

	// for _, kv := range resp.Kvs {
	// 	fmt.Printf("Key: %s, Value: %s\n", kv.Key, kv.Value)
	// }
}
