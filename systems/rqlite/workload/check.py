import pyrqlite.dbapi2 as dbapi2

def check_inserted_data(cursor, start, end):
    # 逐个检查是否插入成功
    for i in range(start, end + 1):
        sql = "SELECT id FROM t WHERE id = ?"
        cursor.execute(sql, (i,))
        result = cursor.fetchone()
        if result:
            # print(f"Row with id {i} inserted successfully.")
            continue
        else:
            print(f"Error: Row with id {i} not found.")
            exit(-1)

def main():
    # 连接到rqlite数据库
    connection = dbapi2.connect(
        host='172.16.237.100',
        port=2379,
    )

    try:
        with connection.cursor() as cursor:

            # 检查插入的数据
            check_inserted_data(cursor, 1, 100)
    finally:
        # 关闭连接
        connection.close()

if __name__ == "__main__":
    main()
    exit(0)
