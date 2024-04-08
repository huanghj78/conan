# docker exec node1 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gsql -p 21000 -d postgres -r -c "CREATE TABLE t (id INT);"

time docker exec node1 env LD_LIBRARY_PATH=/home/opengauss/openGauss/install/lib /home/opengauss/openGauss/install/bin/gsql -p 21000 -d postgres -r -c "EXPLAIN INSERT INTO t VALUES (1);"