percent=$1
nodeNum=$2
duration=$3

docker exec etcd-node$nodeNum-1 /chaosblade-1.3.0/blade create cpu load --cpu-percent $percent --timeout $duration
