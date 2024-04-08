# Record the start time
start_time=$(date +%s)
echo "Stop nodes"
{ 
  docker stop node1 & 
  docker stop node2 & 
  docker stop node3 & 
  wait 
} 
echo "Remove nodes"
{ 
  docker rm node1 & 
  docker rm node2 & 
  docker rm node3 & 
  wait 
}
echo "Start nodes"
docker run -p 21001:21001 -p 21004:21004 -p 21005:21005 --cap-add=SYS_PTRACE --name node1 --privileged -itd opengauss:conan
docker run -p 31001:31001 -p 31004:31004 -p 31005:31005 --cap-add=SYS_PTRACE --name node2 --privileged -itd opengauss:conan 
docker run -p 41001:41001 -p 41004:41004 -p 41005:41005 --cap-add=SYS_PTRACE --name node3 --privileged -itd opengauss:conan 

echo "Wait for opengauss setup, Sleep for 70s"
sleep 70s
echo "Stop nodes"
{ docker exec node1 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl stop -D /home/opengauss/openGauss/data & 
  docker exec node2 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl stop -D /home/opengauss/openGauss/data & 
  docker exec node3 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl stop -D /home/opengauss/openGauss/data & 
  wait
}
echo "Mkdir /home/opengauss/opengauss"
{ docker exec node1 mkdir /home/opengauss/opengauss & 
  docker exec node2 mkdir /home/opengauss/opengauss & 
  docker exec node3 mkdir /home/opengauss/opengauss & 
  wait
}

echo "Init new opengauss"
{ docker exec node1 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_initdb --nodename=gaussdb1 -w Enmo@123 -D /home/opengauss/opengauss/data/ -c &
  docker exec node2 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_initdb --nodename=gaussdb2 -w Enmo@123 -D /home/opengauss/opengauss/data/ -c &
  docker exec node3 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_initdb --nodename=gaussdb3 -w Enmo@123 -D /home/opengauss/opengauss/data/ -c &
  wait 
}

echo "Modify configure"

{ docker exec node1 sh -c 'echo "port = 21000
dcf_ssl = off
dcf_data_path = '\''/home/opengauss/opengauss/dcf_data'\''
dcf_log_path= '\''/home/opengauss/opengauss/dcf_log'\''
dcf_node_id = 1
dcf_config = '\''[{\"stream_id\":1,\"node_id\":1,\"ip\":\"172.17.0.2\",\"port\":21000,\"role\":\"LEADER\"},{\"stream_id\":1,\"node_id\":2,\"ip\":\"172.17.0.3\",\"port\":21000,\"role\":\"FOLLOWER\"}, {\"stream_id\":1,\"node_id\":3,\"ip\":\"172.17.0.4\",\"port\":21000,\"role\":\"FOLLOWER\"}]'\'' 
replconninfo1 = '\''localhost=172.17.0.2 localport=21001 localheartbeatport=21005 localservice=21004 remotehost=172.17.0.3 remoteport=31001 remoteheartbeatport=31005 remoteservice=31004'\'' 
replconninfo2 = '\''localhost=172.17.0.2 localport=21001 localheartbeatport=21005 localservice=21004 remotehost=172.17.0.4 remoteport=41001 remoteheartbeatport=41005 remoteservice=41004'\''" >> /home/opengauss/opengauss/data/postgresql.conf' &

  docker exec node2 sh -c 'echo "port = 21000
#enable_dcf = on
dcf_ssl = off
dcf_data_path = '\''/home/opengauss/opengauss/dcf_data'\''
dcf_log_path= '\''/home/opengauss/opengauss/dcf_log'\''
dcf_node_id = 2
dcf_config = '\''[{\"stream_id\":1,\"node_id\":1,\"ip\":\"172.17.0.2\",\"port\":21000,\"role\":\"LEADER\"},{\"stream_id\":1,\"node_id\":2,\"ip\":\"172.17.0.3\",\"port\":21000,\"role\":\"FOLLOWER\"}, {\"stream_id\":1,\"node_id\":3,\"ip\":\"172.17.0.4\",\"port\":21000,\"role\":\"FOLLOWER\"}]'\'' 
replconninfo1 = '\''localhost=172.17.0.3 localport=31001 localheartbeatport=31005 localservice=31004 remotehost=172.17.0.2 remoteport=21001 remoteheartbeatport=21005 remoteservice=21004'\'' 
replconninfo2 = '\''localhost=172.17.0.3 localport=31001 localheartbeatport=31005 localservice=31004 remotehost=172.17.0.4 remoteport=41001 remoteheartbeatport=41005 remoteservice=41004'\''" >> /home/opengauss/opengauss/data/postgresql.conf' &

  docker exec node3 sh -c 'echo "port = 21000
#enable_dcf = on
dcf_ssl = off
dcf_data_path = '\''/home/opengauss/opengauss/dcf_data'\''
dcf_log_path= '\''/home/opengauss/opengauss/dcf_log'\''
dcf_node_id = 3
dcf_config = '\''[{\"stream_id\":1,\"node_id\":1,\"ip\":\"172.17.0.2\",\"port\":21000,\"role\":\"LEADER\"},{\"stream_id\":1,\"node_id\":2,\"ip\":\"172.17.0.3\",\"port\":21000,\"role\":\"FOLLOWER\"}, {\"stream_id\":1,\"node_id\":3,\"ip\":\"172.17.0.4\",\"port\":21000,\"role\":\"FOLLOWER\"}]'\'' 
replconninfo1 = '\''localhost=172.17.0.4 localport=41001 localheartbeatport=41005 localservice=41004 remotehost=172.17.0.2 remoteport=21001 remoteheartbeatport=21005 remoteservice=21004'\'' 
replconninfo2 = '\''localhost=172.17.0.4 localport=41001 localheartbeatport=41005 localservice=41004 remotehost=172.17.0.3 remoteport=31001 remoteheartbeatport=31005 remoteservice=31004'\''" >> /home/opengauss/opengauss/data/postgresql.conf' &

  wait
}

{ docker exec node1 sh -c 'echo "host all all 0.0.0.0/0 trust" >> /home/opengauss/opengauss/data/pg_hba.conf' &
  docker exec node2 sh -c 'echo "host all all 0.0.0.0/0 trust" >> /home/opengauss/opengauss/data/pg_hba.conf' &
  docker exec node3 sh -c 'echo "host all all 0.0.0.0/0 trust" >> /home/opengauss/opengauss/data/pg_hba.conf' &
  wait
}


{ docker exec node1 mkdir /home/opengauss/opengauss/dcf_data &
  docker exec node1 mkdir /home/opengauss/opengauss/dcf_log &
  docker exec node2 mkdir /home/opengauss/opengauss/dcf_data &
  docker exec node2 mkdir /home/opengauss/opengauss/dcf_log &
  docker exec node3 mkdir /home/opengauss/opengauss/dcf_data &
  docker exec node3 mkdir /home/opengauss/opengauss/dcf_log &
  wait
}


echo "Start all node in standby mode"
{ docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node1 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl start -D /home/opengauss/opengauss/data -M standby &
  docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node2 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl start -D /home/opengauss/opengauss/data -M standby &
  docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node3 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl start -D /home/opengauss/opengauss/data -M standby &
  wait
}
# docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node1 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl start -D /home/opengauss/opengauss/data -M standby 
# docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node2 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl start -D /home/opengauss/opengauss/data -M standby 
# docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node3 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl start -D /home/opengauss/opengauss/data -M standby 

sleep 10s
echo "set node1's run mode"
docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node1  env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl setrunmode -D /home/opengauss/opengauss/data  -v 1 -x minority
sleep 10s
echo "build node2 node3"
# docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node2 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl build -b full -D /home/opengauss/opengauss/data
# docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node3 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl build -b full -D /home/opengauss/opengauss/data
{ docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node2 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl build -b full -D /home/opengauss/opengauss/data &
  docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node3 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl build -b full -D /home/opengauss/opengauss/data &
  wait 
}
sleep 30s
echo "set node1's run mode"
docker exec  -e GAUSSHOME=/home/opengauss/openGauss/install node1 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl setrunmode -D /home/opengauss/opengauss/data -x normal
# exit
echo "Copy file to docker"
{ docker cp /root/cpfi-v2/client/client_c/libdcfi.so node1:/home/opengauss/openGauss/install/lib/libdcfi.so &
  docker cp /root/cpfi-v2/systems/opengauss/new_so/libdcf.so node1:/home/opengauss/openGauss/install/lib/libdcf.so &

  docker cp /root/cpfi-v2/client/client_c/libdcfi.so node2:/home/opengauss/openGauss/install/lib/libdcfi.so &
  docker cp /root/cpfi-v2/systems/opengauss/new_so/libdcf.so node2:/home/opengauss/openGauss/install/lib/libdcf.so &

  docker cp /root/cpfi-v2/client/client_c/libdcfi.so node3:/home/opengauss/openGauss/install/lib/libdcfi.so &
  docker cp /root/cpfi-v2/systems/opengauss/new_so/libdcf.so node3:/home/opengauss/openGauss/install/lib/libdcf.so &
  wait
}

# docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node1 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl stop -D /home/opengauss/opengauss/data 
# docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node1 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl start -D /home/opengauss/opengauss/data -M standby 
# # sleep 10s
# docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node2 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl stop -D /home/opengauss/opengauss/data 
# docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node2 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl start -D /home/opengauss/opengauss/data -M standby 
# # sleep 10s
# docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node3 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl stop -D /home/opengauss/opengauss/data 
# docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node3 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl start -D /home/opengauss/opengauss/data -M standby 

{
  docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node1 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl stop -D /home/opengauss/opengauss/data &
  docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node2 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl stop -D /home/opengauss/opengauss/data &
  docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node3 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl stop -D /home/opengauss/opengauss/data &
  wait 
}

{ 
  docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node1 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl start -D /home/opengauss/opengauss/data -M standby &
  docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node2 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl start -D /home/opengauss/opengauss/data -M standby &
  docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node3 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl start -D /home/opengauss/opengauss/data -M standby &
  wait
}

{
  docker exec node1 sudo  yum install -y iproute-tc &
  docker exec node2 sudo  yum install -y iproute-tc &
  docker exec node3 sudo  yum install -y iproute-tc &
  wait
}

{
  docker cp /root/cpfi-v2/scripts/chaosblade-1.3.0 node1:/home/opengauss &
  docker cp /root/cpfi-v2/scripts/chaosblade-1.3.0 node2:/home/opengauss &
  docker cp /root/cpfi-v2/scripts/chaosblade-1.3.0 node3:/home/opengauss &
  wait
}

# query status
{ docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node1 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl query -D /home/opengauss/opengauss/data &
  docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node2 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl query -D /home/opengauss/opengauss/data &
  docker exec -e GAUSSHOME=/home/opengauss/openGauss/install node3 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gs_ctl query -D /home/opengauss/opengauss/data &
  wait
}
end_time=$(date +%s)
execution_time=$((end_time - start_time))
echo "Total execution time: ${execution_time} seconds."
echo "===================END==================="