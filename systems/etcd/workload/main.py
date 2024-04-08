import etcd3
import os

etcd_host = "172.16.238.100"
etcd_port = 2379

etcd = etcd3.client(host=etcd_host, port=etcd_port)

# 设置键值对
def set_key_value(key, value):
    etcd.put(key, value)
    # print(f"Set key: {key}, value: {value}")

# 获取键的值
def get_key_value(key):
    result, _ = etcd.get(key)
    return result
    # if result:
    #     print(f"Key: {key}, Value: {result.decode('utf-8')}")
    # else:
    #     print(f"Key: {key} not found")

def del_key_value(key):
    etcd.delete(key)

# 监视键的更改
def watch_key(key):
    events, cancel = etcd.watch(key)
    for event in events:
        print(f"Key: {event.key}, Value: {event.value}")

if __name__ == "__main__":
    for i in range(2):
        key = f"k{i}"
        value = f"v{i}"
        set_key_value(key, value)

    