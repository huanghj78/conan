docker-compose  down -v
docker-compose  up -d


{
  docker cp ../../scripts/chaosblade-1.3.0 rqlite-node1-1:/ &
  docker cp ../../scripts/chaosblade-1.3.0 rqlite-node2-1:/ &
  docker cp ../../scripts/chaosblade-1.3.0 rqlite-node3-1:/ &
  wait
}

sleep 5
