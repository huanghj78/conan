import requests
import json

def get_leader_node_id(rqlite_url):
    try:
        # 发送HTTP请求获取rqlite状态
        response = requests.get(f"{rqlite_url}/status")
        response.raise_for_status()  # 检查请求是否成功

        # 解析JSON数据
        data = response.json()

        # 提取leader节点编号
        leader_node_id = data["store"]["leader"]["node_id"]

        return leader_node_id
    except requests.exceptions.RequestException as e:
        print(f"Error: {e}")
        return None

if __name__ == "__main__":
    rqlite_url = "http://172.16.237.100:2379"

    leader_node_id = get_leader_node_id(rqlite_url)

    if leader_node_id is not None:
        print(f"{leader_node_id}")
        exit(0)
    else:
        print("Failed to retrieve leader node ID.")
        exit(-1)
