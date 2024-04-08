delayTime=$1
nodeNum=$2
duration=$3

../../scripts/chaosblade-1.3.0/blade create docker network delay --time $delayTime --timeout $duration --interface eth0 --local-port 2380 --container-name rqlite-node$nodeNum-1