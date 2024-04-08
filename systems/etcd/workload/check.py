import subprocess
import shlex
import time
import signal
from subprocess import TimeoutExpired

# 设置超时时间（单位：秒）
timeout = 5
num = 100


def run_shell_command(command, timeout):
    try:
        cmd_args = shlex.split(command)
        process = subprocess.Popen(cmd_args, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        stdout, stderr = process.communicate(timeout=timeout)
        return process.returncode, stdout, stderr

    except TimeoutExpired:
        process.terminate()
        return -1, b'', b'Command timed out'



for i in range(num):
    v = f"v{i}"
    ret_code, stdout, stderr = run_shell_command(f"../etcdctl  --endpoints=http://172.16.238.100:2379 get k{i}", timeout)
    if ret_code == 0:
        if len(stdout.decode().split('\n')) != 3:
            print("No Value")
            exit(1)
        value = stdout.decode().split('\n')[1]
        if v != value:
            print(f"Expect {v}, get {value}")
            exit(1)
    elif ret_code == -1:
        print("Dail node1 timeout")
        exit(1)
    else:
        print(stderr.decode())
        exit(1)
    

for i in range(num):
    v = f"v{i}"
    ret_code, stdout, stderr = run_shell_command(f"etcdctl  --endpoints=http://172.16.238.101:2379 get k{i}", timeout)
    if ret_code == 0:
        if len(stdout.decode().split('\n')) != 3:
            print("No Value")
            exit(1)
        value = stdout.decode().split('\n')[1]
        if v != value:
            print(f"Expect {v}, get {value}")
            exit(1)
    elif ret_code == -1:
        print("Dail node2 timeout")
        exit(1)
    else:
        print(stderr.decode())
        exit(1)

for i in range(num):
    v = f"v{i}"
    ret_code, stdout, stderr = run_shell_command(f"etcdctl  --endpoints=http://172.16.238.102:2379 get k{i}", timeout)
    if ret_code == 0:
        if len(stdout.decode().split('\n')) != 3:
            print("No Value")
            exit(1)
        value = stdout.decode().split('\n')[1]
        if v != value:
            print(f"Expect {v}, get {value}")
            exit(1)
    elif ret_code == -1:
        print("Dail node3 timeout")
        exit(1)
    else:
        print(stderr.decode())
        exit(1)

exit(0)