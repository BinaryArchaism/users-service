CREATE TABLE users
(
    FirstName String,
    LastName String,
    Email String,
    Age UInt32
) ENGINE = Kafka('localhost:9092', 'topic', 'clickhouse', 'JSONEachRow');