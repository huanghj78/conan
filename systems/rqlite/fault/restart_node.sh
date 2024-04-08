nodeNum=$1
docker-compose -f $CONAN_PATH/systems/rqlite/docker-compose.yaml restart node$nodeNum
