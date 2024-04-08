docker-compose  down -v
docker-compose up -d


{
  docker cp ../../scripts/chaosblade-1.3.0 etcd-node1-1:/ &
  docker cp ../../scripts/chaosblade-1.3.0 etcd-node2-1:/ &
  docker cp ../../scripts/chaosblade-1.3.0 etcd-node3-1:/ &
  wait
}

sleep 5
