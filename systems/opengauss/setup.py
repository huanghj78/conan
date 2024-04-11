import os
path = os.environ.get("CONAN_PATH")

os.system(f"{path}/systems/opengauss/setup.sh")

