import pyrqlite.dbapi2 as dbapi2
from datetime import datetime, time
def create_table(cursor):
    # 删除表t
    cursor.execute('DROP TABLE IF EXISTS t')
    
    # 创建表t
    cursor.execute('CREATE TABLE t (id INT)')

def insert_data(cursor, start, end):
    # 插入从start到end的数字
    data = [(i,) for i in range(start, end + 1)]
    cursor.executemany('INSERT INTO t (id) VALUES (?)', seq_of_parameters=data)

def check_inserted_data(cursor, start, end):
    # 逐个检查是否插入成功
    for i in range(start, end + 1):
        sql = "SELECT id FROM t WHERE id = ?"
        cursor.execute(sql, (i,))
        result = cursor.fetchone()
        if result:
            print(f"Row with id {i} inserted successfully.")
        else:
            print(f"Error: Row with id {i} not found.")

def main():
    # 连接到rqlite数据库
    connection = dbapi2.connect(
        host='172.16.237.100',
        port=2379,
    )

    try:
        with connection.cursor() as cursor:
            # 创建表t
            create_table(cursor)
            
            # 插入数据
            insert_data(cursor, 1, 1000)

            # 检查插入的数据
            # check_inserted_data(cursor, 1, 100)
    except Exception as e:
        print(f"Error: {e}")
        exit(-1)
    finally:
        connection.close()

if __name__ == "__main__":
    begin_time = datetime.now().time()
    begin_time_with_ms = begin_time.strftime('%H:%M:%S.%f')
    print(begin_time_with_ms)
    main()
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
