# Docker cac cau lenh can thiet

```bash
# Start docker with port
docker run -ti -p 15000:15000 -p 15100:15100 -p 15101:15101 -p 15102:15102 -p 15103:15103 -p 15104:15104 -p 17100:17100 -p 17101:17101 -p 17102:17102 -p 17103:17103 -p 17104:17104 vitess/base bash

docker run -t -d -p 15000:15000 -p 15100:15100 -p 15101:15101 -p 15102:15102 -p 15103:15103 -p 15104:15104 -p 17100:17100 -p 17101:17101 -p 17102:17102 -p 17103:17103 -p 17104:17104 vitess/base

# Connect to Docker
docker exec -it c925326640a9 bash

# Copy
docker cp ./client.py 04c613302de1:/vt/src/vitess.io/vitess/examples/local/

docker cp ./test_client_1.go 04c613302de1:/vt/src/vitess.io/vitess/examples/local/
docker cp ./create_like_ratings_table.sql 04c613302de1:/vt/src/vitess.io/vitess/examples/local/

docker cp ./vschema.json 04c613302de1:/vt/src/vitess.io/vitess/examples/local/
```

# Thuc thi cac cau lenh can thiet

> Cac cau lenh nay neu chay qua nhanh se gay ra loi vi no la he thong cluster nen no khong the dong bo kip thoi

```bash
cd examples/local/
./zk-up.sh
sleep 5
./vtctld-up.sh
sleep 5
./vttablet-up.sh
sleep 5
./lvtctl.sh InitShardMaster -force test_keyspace/0 test-100
sleep 5
./lvtctl.sh ListAllTablets test
sleep 5
./lvtctl.sh ApplySchema -sql "$(cat create_test_table.sql)" test_keyspace
sleep 5
  ./lvtctl.sh RebuildVSchemaGraph
sleep 5
./vtgate-up.sh
sleep 5
./client.sh
sleep 5
```


# Test gia dinh ve ve viec insert lien tuc
```bash
while true; do ./client.sh; sleep 1; done
```


# Others

```bash
# Chu y la phai co cai keyspace:
./lvtctl.sh ApplySchema -sql "$(cat create_like_ratings_table.sql)" test_keyspace

./lvtctl.sh ApplySchema -sql "$(cat create_product_table.sql)" test_keyspace

./lvtctl.sh ApplySchema -sql "ALTER TABLE messages ADD COLUMN `a_test_column` SMALLINT(6) NOT NULL DEFAULT 0 AFTER `message`" test_keyspace


./lvtctl.sh ApplyVSchema -vschema "$(cat vschema.json)" test_keyspace

./lvtctl.sh ListAllTablets test

# ! KHONG work
# ./lvtctl.sh VtGateExecute -enable_queries  -sql "SELECT * FROM messages LIMIT 10"

while true; do go run test_client_1.go; sleep 1; done

nohup sh -c 'while true; do go run test_client_1.go; sleep 1; done >log1.txt' &
nohup sh -c 'while true; do go run test_client_1.go; sleep 1; done >log2.txt' &


docker cp ./vschema.json bd9eaf616413:/vt/src/vitess.io/vitess/examples/local/
docker cp ./create_like_ratings_table.sql c925326640a9:/vt/src/vitess.io/vitess/examples/local/

```

```bash
# Cau lenh work
## Cau lenh connect vo db
mysql --user=vt_appdebug --host=localhost --socket=/vt/vtdataroot/vt_0000000100/mysql.sock --port=17100

mysql --port=15306 --host=localhost --socket=/tmp/mysql.sock --user=vt_appdebug

mysql --port=15306 --host=localhost --socket=/tmp/mysql.sock --user=mysql_user3 --default-auth=mysql_native_password

./lvtctl.sh ExecuteFetchAsDba test-0000000100 "select * from like_ratings where id = 100000"

./lvtctl.sh ExecuteFetchAsDba test-0000000100 "update like_ratings set is_delete = 1 where id = 100000"

./lvtctl.sh ExecuteFetchAsDba test-0000000103 "update like_ratings set is_delete = 1 where id = 100000"

go run test_client_1.go --incr 400000

## Cau lenh mysql
show databases;
use mysql;
show tables;
select * from user as u where u.user = 'vt_appdebug' \G;

select Host, User, Select_priv as SEL, Insert_priv as INS, Update_priv as UP, Delete_priv as DEL, Create_priv as `CRE`, Drop_priv as `DR`, Reload_priv as `REL`, Shutdown_priv as `SD`, Process_priv as `PRO`, File_priv as `File`, Grant_priv as `Grant`, References_priv as `Ref`, Index_priv as `IDX`, Alter_priv as `ALT`, Show_db_priv as SHOWDB, Super_priv as `Super`, Create_tmp_table_priv as CRETmpTbl, Lock_tables_priv as `Lock`, Execute_priv as `Exec`, Repl_slave_priv as `Relp`, Create_view_priv as CREView, Show_view_priv as SView, Create_routine_priv as CRERT, Alter_routine_priv as  ALTRT, Create_user_priv as CREUser, Event_priv as `Event`, Trigger_priv as `Trigger`, Create_tablespace_priv as CRETblSpc  from user as u where u.user LIKE 'vt%'


ALTER TABLE messages
ADD COLUMN `a_test_column` SMALLINT(6) NOT NULL DEFAULT 0 AFTER `message`
```



```bash
# Cac cau lenh khong work
mysql --user=mysql_user --host=127.0.0.1 --socket=/vt/vtdataroot/vt_0000000100/mysql.sock --port=17100

mysql --user=mysql_user --password=mysql_password --host=localhost --socket=/vt/vtdataroot/vt_0000000100/mysql.sock --port=17100

mysql --user=mysql_user --password=mysql_password --host=c925326640a9 --socket=/vt/vtdataroot/vt_0000000100/mysql.sock --port=17100

mysql --user=mysql_user --password=mysql_password --host=172.17.0.2 --socket=/vt/vtdataroot/vt_0000000100/mysql.sock --port=17100


mysql --user=mysql_user --password=mysql_password --host=127.0.0.1 --socket=/vt/vtdataroot/vt_0000000100/mysql.sock --port=17100



mysql --user=vt_appdebug --host=127.0.0.1 --socket=/vt/vtdataroot/vt_0000000100/mysql.sock --port=17100

mysql --user=vt_appdebug --host=localhost --socket=/vt/vtdataroot/vt_0000000100/mysql.sock --port=17100

mysql --user=vt_appdebug --host=c925326640a9 --socket=/vt/vtdataroot/vt_0000000100/mysql.sock --port=17100

mysql --user=vt_appdebug --host=172.17.0.2 --socket=/vt/vtdataroot/vt_0000000100/mysql.sock --port=17100


mysql --user=vt_appdebug --host=127.0.0.1 --socket=/vt/vtdataroot/vt_0000000100/mysql.sock --port=17100


/usr/sbin/mysqld --defaults-file=/vt/vtdataroot/vt_0000000100/my.cnf --basedir=/usr --datadir=/vt/vtdataroot/vt_0000000100/data --plugin-dir=/usr/lib/mysql/plugin --log-error=/vt/vtdataroot/vt_0000000100/error.log --pid-file=/vt/vtdataroot/vt_0000000100/mysql.pid --socket=/vt/vtdataroot/vt_0000000100/mysql.sock --port=17100


/home/phutp/noodle/vitess/go/vt/proto/query/query.pb.go

```