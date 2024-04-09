#!/bin/bash


python_script="$CONAN_PATH/systems/rqlite/findLeader.py"  

python_output=$(python3 "$python_script")

echo "$python_output"
