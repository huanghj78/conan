percent=$1
nodeNum=$2
duration=$3
docker exec node$nodeNum sudo /home/opengauss/chaosblade-1.3.0/blade create cpu fullload --cpu-percent $percent --timeout $duration 