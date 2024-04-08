docker-compose  down -v
docker-compose up -d


{
  docker cp /root/cpfi-v2/scripts/chaosblade-1.3.0 bin-node1-1:/ &
  docker cp /root/cpfi-v2/scripts/chaosblade-1.3.0 bin-node2-1:/ &
  docker cp /root/cpfi-v2/scripts/chaosblade-1.3.0 bin-node3-1:/ &
  wait
}

sleep 5
