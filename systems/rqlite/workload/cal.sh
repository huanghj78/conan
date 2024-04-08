time curl -XPOST '172.16.237.100:2379/db/execute?pretty&timings' -H "Content-Type: application/json" -d "[
    \"INSERT INTO t(id) VALUES('1')\"
]"