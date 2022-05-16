package parser

import "flag"

type Config struct {
	dbPort int
	dbConnectionURL string
	dbJaeger_url string
	dbSentry_url string
	dbKafka_broker string
	dbSome_app_id string
	dbSome_app_key string
}

func MyParser() Config {
	var c Config
	c.dbPort := flag.Int("port", 8080, "port for db")
	c.dbConnectionURL := flag.String("dbConnectionURL", "db_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "connection")
	c.dbJaeger_url := flag.String("jaeger_url", "http://jaeger:16686", "db jaeger_url")
	c.dbSentry_url := flag.String("sentry_url", "http://sentry:9000", "db sentry_url")
	c.dbKafka_broker := flag.String("kafka_url", "kafka:9092", "db kafka_url")
	c.dbSome_app_id := flag.String("some_app_id", "testid", "db some_app_id")
	c.dbSome_app_key := flag.String("some_app_key", "testkey", "db some_app_key")
	return c
}
