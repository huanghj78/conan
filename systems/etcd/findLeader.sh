#!/bin/bash

# 定义etcd节点列表
endpoints=("http://172.16.238.100:2379" "http://172.16.238.101:2379" "http://172.16.238.102:2379")

# 初始化leader序号为0
leader_index=0

# 初始化leader计数器
leader_count=0

# 循环遍历每个节点
for ((i=0; i<${#endpoints[@]}; i++)); do
    # 执行etcdctl命令，获取节点信息
    result=$($CONAN_PATH/systems/etcd/etcdctl --endpoints=${endpoints[$i]} endpoint status)

    # 提取第五个参数，判断是否为leader节点
    is_leader=$(echo $result | awk '{print $6}')

    # 如果是leader节点，更新leader序号并增加计数器
    if [ "$is_leader" == "true," ]; then
        leader_index=$((i+1))
        leader_count=$((leader_count+1))
    fi
done

# 判断leader节点数量
if [ $leader_count -eq 1 ]; then
    # 打印leader节点序号
    echo "$leader_index"
elif [ $leader_count -eq 0 ]; then
    # 没有找到leader节点
    echo "没有找到Leader节点"
    exit -1
else
    # 发现多个leader节点
    echo "发现多个Leader节点"
    leader_index=-1
    exit -1
fi

# 返回leader节点序号或-1
exit 0
