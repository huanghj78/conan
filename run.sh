system=$1
mode=$2
jq --arg mode "$mode" '.mode = $mode' ./config/$system.json > tmp && mv tmp ./config/$system.json
export CONAN_PATH=$(pwd)

cd server && ./conan-server ../config/$system.json
