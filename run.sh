system=$1

export CONAN_PATH=$(pwd)

cd server && ./conan-server ../config/$system.json
