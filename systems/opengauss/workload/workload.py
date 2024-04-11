import subprocess
import shlex
import time
import signal
from subprocess import TimeoutExpired
import os
# 设置超时时间（单位：秒）
timeout = 60
num = 10
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



import datetime

current_time = datetime.datetime.now().time()
current_time_with_ms = current_time.strftime('%H:%M:%S.%f')

print(current_time_with_ms)

# time.sleep(5)

ret_code, stdout, stderr = run_shell_command(f"{path}/systems/opengauss/workload/workload.sh {num}", timeout)
print(stdout)
exit(ret_code)