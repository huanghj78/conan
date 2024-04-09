import os

os.system("rm -r /root/rqlite/raft")
os.system("cp -r /root/Conan/systems/rqlite/raft /root/rqlite/raft")
os.system("cp -r /root/Conan/client /root/rqlite/raft/rafi.client")
os.system("rm /root/rqlite/raft/rafi.client/go.mod")
# 已修改好rqlite下的go.mod
os.system("cd /root/rqlite; go install ./...")

# os.system("cd /root/.gvm/pkgsets/go1.21/global/bin/")
# os.system("docker build -t rqlited:latest . --no-cache")