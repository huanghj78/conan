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
primary_node=""
# Loop through the nodes
for node in "${nodes[@]}"; do
    # Execute gsql command on the current node
    result=$(docker exec "$node" env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gsql -p 21000 -d postgres -r -c "SELECT pg_is_in_recovery();"| grep -oE '[tf]')
    echo $result
    # Check the result and print the node's status
    if [ "$result" == "f" ]; then
        primary_node="$node"
        echo "Node $node is the primary (master) node."
    elif [ "$result" == "t" ]; then
        continue
        echo "Node $node is a standby (replica) node."
    else
        echo "Error while checking node $node."
        exit 1
    fi
done

# Check if a primary node was found
if [ -n "$primary_node" ]; then
    # Execute commands on the primary node
    docker exec "$primary_node" env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gsql -p 21000 -d postgres -r -c "CREATE TABLE t (id INT);"

    # Loop to insert values into the table
    for ((i = 1; i <= num; i++)); do
        docker exec "$primary_node" env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gsql -p 21000 -d postgres -r -c "INSERT INTO t VALUES ($i);"
        echo "Inserted value $i into table t."
    done
else
    exit 1
    echo "No primary node found."
fi

exit 0