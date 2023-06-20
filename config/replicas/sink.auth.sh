#/bin/bash

curl --silent --show-error --fail -X POST http://connect:8083/connectors \
-H  "Content-Type:application/json" --data '{"name":"auth-mysql-sink","config":{"connector.class":"io.confluent.connect.jdbc.JdbcSinkConnector","tasks.max":"1","topics":"Users","connection.url":"${file:/data/credentials.properties:AUTH_DB_CONNECTION_STRING}","transforms":"unwrap","transforms.unwrap.type":"io.debezium.transforms.ExtractNewRecordState","transforms.unwrap.drop.tombstones":"false","auto.create":"true","insert.mode":"upsert","delete.enabled":"true","pk.fields":"Id","pk.mode":"record_key", "max.retries":10,"retry.backoff.ms":10000}}'

# {
#   "name": "auth-mysql-sink",
#   "config": {
#     "connector.class": "io.confluent.connect.jdbc.JdbcSinkConnector",
#     "tasks.max": "1",
#     "topics": "Users",
#     "connection.url": "jdbc:mysql://root:secret@auth-mysql:3306/UserMicroservice",
#     "transforms": "unwrap",
#     "transforms.unwrap.type": "io.debezium.transforms.ExtractNewRecordState",
#     "transforms.unwrap.drop.tombstones": "false",
#     "auto.create": "true",
#     "insert.mode": "upsert",
#     "delete.enabled": "true",
#     "pk.fields": "id",
#     "pk.mode": "record_key"
#   }
# }
