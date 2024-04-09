mkdir logs

docker pull 13822030416/etcd:conan
docker pull 13822030416/rqlite:conan
docker pull 13822030416/opengauss:conan

docker tag 13822030416/etcd:conan etcd:conan
docker tag 13822030416/rqlite:conan rqlite:conan
docker tag 13822030416/opengauss:conan opengauss:conan

apt-get install jq