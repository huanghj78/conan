#!/bin/bash

# Check if num_iterations is provided as a command-line argument
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <num_iterations>"
    exit 1
fi

# Assign the command-line argument to num_iterations
num="$1"

# Define an array of node names
nodes=("node1" "node2" "node3")

# Loop through the nodes to execute SELECT query
for node in "${nodes[@]}"; do
    # Execute the SELECT query and store the result in an array
    result=($(docker exec "$node" env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gsql -p 21000 -d postgres -r -c "SELECT * FROM t;" | grep -v "(.*rows)" |tail -n +3))
    echo "${result[@]}"
    # Compare the result with the expected values
    expected_result=()
    for ((i = 1; i <= num; i++)); do
        expected_result+=("$i")
    done
    # echo "${expected_result[@]}"
    # Check if the arrays match
    if [ "${result[*]}" == "${expected_result[*]}" ]; then
        continue
        # echo "Values on $node match the inserted values."
    else
        # echo "Values on $node do not match the inserted values."
        exit 1
    fi
done

exit 0

