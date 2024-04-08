#!/bin/bash
nodes=("node1" "node2" "node3")
primary_node=""
primary_node_number=""
# Loop through the nodes
for node in "${nodes[@]}"; do
    # Execute gsql command on the current node
    result=$(docker exec "$node" env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gsql -p 21000 -d postgres -r -c "SELECT pg_is_in_recovery();"| grep -oE '[tf]')
    # echo $result
    
    # Check the result and print the node's status
    if [ "$result" == "f" ]; then
        primary_node="$node"
        # Extract the node number from the node name (assuming the format is "nodeX")
        primary_node_number="${node#"node"}"
        # echo "Node $node is the primary (master) node."
        break  # Stop the loop once the primary node is found
    fi
done

# Output the primary node's number
echo $primary_node_number
exit 0
