# projectPath=$1
docker-compose -f $CONAN_PATH/systems/etcd/docker-compose.yaml down -v
docker-compose -f $CONAN_PATH/systems/etcd/docker-compose.yaml up -d


{
  docker cp $CONAN_PATH/scripts/chaosblade-1.3.0 etcd-node1-1:/ &
  docker cp $CONAN_PATH/scripts/chaosblade-1.3.0 etcd-node2-1:/ &
  docker cp $CONAN_PATH/scripts/chaosblade-1.3.0 etcd-node3-1:/ &
  wait
}

sleep 5
