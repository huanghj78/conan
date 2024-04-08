import subprocess
import shlex
import os
from datetime import datetime, time
import signal
from subprocess import TimeoutExpired

# 设置超时时间（单位：秒）
timeout = 30
num = 1

path = os.environ.get("CONAN_PATH")
def run_shell_command(command, timeout):
    try:
        cmd_args = shlex.split(command)
        process = subprocess.Popen(cmd_args, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        stdout, stderr = process.communicate(timeout=timeout)
        return process.returncode, stdout, stderr

    except TimeoutExpired:
        process.terminate()
        return -1, b'', b'Command timed out'



begin_time = datetime.now().time()
begin_time_with_ms = begin_time.strftime('%H:%M:%S.%f')

print(begin_time_with_ms)

# time.sleep(5)
for i in range(num):
    ret_code, stdout, stderr = run_shell_command(f"{path}/systems/etcd/etcdctl --dial-timeout=5s  --endpoints=http://172.16.238.100:2379,http://172.16.238.101:2379,http://172.16.238.102:2379 put k{i} v{i}", timeout)
    if ret_code == 0:
        if stdout.decode().split('\n')[0] != "OK":
            print("Put error")
            exit(1)
    elif ret_code == -1:
        print("Dail timeout")
        exit(1)
    else:
        print(stderr.decode())
        current_time = datetime.now().time()
        current_time = datetime.now().time()
        current_time_with_ms = current_time.strftime('%H:%M:%S.%f')

        print(current_time_with_ms)
        exit(1)
end_time = datetime.now().time()
end_time_with_ms = end_time.strftime('%H:%M:%S.%f')
print(end_time_with_ms)
current_date = datetime.now().date()
datetime1 = datetime.combine(current_date, end_time)
datetime2 = datetime.combine(current_date, begin_time)

# 计算时间差
time_difference = datetime1 - datetime2

# 输出时间差
print("时间差为：", time_difference)
exit(0)