percent=$1
nodeNum=$2
duration=$3
$CONAN_PATH/scripts/chaosblade-1.3.0/blade create docker network loss --percent $percent --timeout $duration --interface eth0 --local-port 2380 --container-name rqlite-node$nodeNum-1
