nodeNum=$1
docker-compose -f $CONAN_PATH/systems/etcd/docker-compose.yaml restart node$nodeNum