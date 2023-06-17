#/bin/bash

# combination of --silent --show-error --faild make cURL fail with exitCode
# other than 0 in case of failed request
# this is needed to container restart on failure

curl --silent --show-error --fail -X POST http://connect:8083/connectors \
-H  "Content-Type:application/json" --data '{"name":"user-mysql-master","config":{"connector.class":"io.debezium.connector.mysql.MySqlConnector","tasks.max":"1","topic.prefix":"user-main","database.hostname":"${file:/data/credentials.properties:DB_HOST}","database.port":"${file:/data/credentials.properties:DB_PORT}","database.user":"${file:/data/credentials.properties:DB_USER}","database.password":"${file:/data/credentials.properties:DB_PASSWORD}","database.server.id":"1","database.dbname":"${file:/data/credentials.properties:DB_NAME}","database.server.name":"dbserver1","database.include.list":"UserMicroservice","database.allowPublicKeyRetrieval":"true","database.history.kafka.bootstrap.servers":"kafka:9092","database.history.kafka.topic":"schema-changes.UserMicroservice","schema.history.internal.kafka.bootstrap.servers":"kafka:9092","schema.history.internal.kafka.topic":"schema-changes.UserMicroservice","include.schema.changes":"true","transforms":"route","transforms.route.type":"org.apache.kafka.connect.transforms.RegexRouter","transforms.route.regex":"([^.]+)\\.([^.]+)\\.([^.]+)","transforms.route.replacement":"$3", "max.retries":10,"retry.backoff.ms":10000}}' 

# {
#   "name": "user-mysql-master",
#   "config": {
#     "connector.class": "io.debezium.connector.mysql.MySqlConnector",
#     "tasks.max": "1",
#     "topic.prefix": "main",
#     "database.hostname": "${file:/data/credentials.properties:DB_HOST}",
#     "database.port": "${file:/data/credentials.properties:DB_PORT}",
#     "database.user": "${file:/data/credentials.properties:DB_USER}",
#     "database.password": "${file:/data/credentials.properties:DB_PASSWORD}",
#     "database.server.id": "1",
#     "database.dbname": "${file:/data/credentials.properties:DB_NAME}",
#     "database.server.name": "dbserver1",
#     "database.include.list": "UserMicroservice",
#     "database.allowPublicKeyRetrieval": "true",
#     "database.history.kafka.bootstrap.servers": "kafka:9092",
#     "database.history.kafka.topic": "schema-changes.UserMicroservice",
#     "schema.history.internal.kafka.bootstrap.servers": "kafka:9092",
#     "schema.history.internal.kafka.topic": "schema-changes.UserMicroservice",
#     "include.schema.changes": "true",
#     "transforms": "route",
#     "transforms.route.type": "org.apache.kafka.connect.transforms.RegexRouter",
#     "transforms.route.regex": "([^.]+)\\.([^.]+)\\.([^.]+)",
#     "transforms.route.replacement": "$3",
#     "max.retries": 10,
#     "retry.backoff.ms": 10000
#   }
# }