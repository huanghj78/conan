percent=$1
nodeNum=$2
duration=$3
num=$((nodeNum + 1))
docker exec node$nodeNum sudo /home/opengauss/chaosblade-1.3.0/blade create network loss --percent $percent  --interface eth0 --local-port 21000,${num}1001,${num}1004,${num}1005  --timeout $duration