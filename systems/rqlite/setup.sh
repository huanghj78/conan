docker-compose -f /root/cpfi-v2/systems/rqlite/docker-compose.yaml down -v
docker-compose -f /root/cpfi-v2/systems/rqlite/docker-compose.yaml up -d


{
  docker cp /root/cpfi-v2/scripts/chaosblade-1.3.0 rqlite-node1-1:/ &
  docker cp /root/cpfi-v2/scripts/chaosblade-1.3.0 rqlite-node2-1:/ &
  docker cp /root/cpfi-v2/scripts/chaosblade-1.3.0 rqlite-node3-1:/ &
  wait
}

sleep 5
