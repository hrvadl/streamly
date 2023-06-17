FROM confluentinc/cp-kafka-connect:latest

#If you want to run a local build of the connector, uncomment the COPY command and make sure the JAR file is in the directory path
#COPY mongo-kafka-connect-<<INSERT BUILD HERE>>3-all.jar /usr/share/confluent-hub-components
RUN confluent-hub install --no-prompt --verbose confluentinc/kafka-connect-jdbc:10.7.2
RUN confluent-hub install --no-prompt --verbose mongodb/kafka-connect-mongodb:latest
RUN confluent-hub install --no-prompt --verbose debezium/debezium-connector-mysql:latest
