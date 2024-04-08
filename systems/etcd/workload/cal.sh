# #!/bin/bash

# # 设置 etcdctl 命令
# ETCDCTL_CMD="etcdctl --dial-timeout=5s --endpoints=http://172.16.238.100:2379,http://172.16.238.101:2379,http://172.16.238.102:2379 put k1 v1"

# # 设置执行次数
# NUM_RUNS=100

# # 初始化总执行时间
# total_time=0

# # 循环执行命令
# for ((i=1; i<=$NUM_RUNS; i++))
# do
#     # 记录开始时间
#     start_time=$(date +%s.%N)

#     # 执行命令
#     $ETCDCTL_CMD

#     # 记录结束时间
#     end_time=$(date +%s.%N)

#     # 计算本次执行时间并累加到总执行时间
#     elapsed_time=$(echo "$end_time - $start_time" | bc)
#     total_time=$(echo "$total_time + $elapsed_time" | bc)
# done

# # 计算平均执行时间
# average_time=$(echo "scale=6; $total_time / $NUM_RUNS" | bc)

# # 输出结果
# echo "Total time: $total_time seconds"
# echo "Average time per command: $average_time seconds"

#!/bin/bash

# 设置 etcdctl 命令
ETCDCTL_CMD="etcdctl --dial-timeout=5s --endpoints=http://172.16.238.100:2379,http://172.16.238.101:2379,http://172.16.238.102:2379 put k1 v1"

time etcdctl --dial-timeout=5s --endpoints=http://172.16.238.100:2379,http://172.16.238.101:2379,http://172.16.238.102:2379 put k1 v1

# # 设置执行次数
# NUM_RUNS=100

# # 初始化总执行时间
# total_time=0

# # 循环执行命令
# for ((i=1; i<=$NUM_RUNS; i++))
# do
#     # 记录开始时间
#     start_time=$(date +%s.%N)

#     # 执行命令
#     $ETCDCTL_CMD

#     # 记录结束时间
#     end_time=$(date +%s.%N)

#     # 计算本次执行时间并累加到总执行时间
#     elapsed_time=$(echo "$end_time - $start_time" | bc)
#     total_time=$(echo "$total_time + $elapsed_time" | bc)
# done

# # 计算平均执行时间
# average_time=$(echo "scale=6; $total_time / $NUM_RUNS" | bc)

# # 输出结果
# echo "Total time: $total_time seconds"
# echo "Average time per command: $average_time seconds"

