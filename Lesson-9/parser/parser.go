package parser

import (
	"encoding/json"
)

type Config struct {
	dbPort          string `json:"dbPort"`
	dbConnectionURL string `json:"dbConnectionURL"`
	dbJaeger_url    string `json:"dbJaeger_url"`
	dbSentry_url    string `json:"dbSentry_url"`
	dbKafka_broker  string `json:"dbKafka_broker"`
	dbSome_app_id   string `json:"dbSome_app_id"`
	dbSome_app_key  string `json:"dbSome_app_key"`
}

func p_json() {
	var config2 Config
	data, err := os.Readfile("db.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &config2); err != nil {
		panic(err)
	}
	return config2
}

//func MyParser() Config {
//	var c Config
//	c.dbPort := flag.Int("port", 8080, "port for db")
//	c.dbConnectionURL := flag.String("dbConnectionURL", "db_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "connection")
//	c.dbJaeger_url := flag.String("jaeger_url", "http://jaeger:16686", "db jaeger_url")
//	c.dbSentry_url := flag.String("sentry_url", "http://sentry:9000", "db sentry_url")
//	c.dbKafka_broker := flag.String("kafka_url", "kafka:9092", "db kafka_url")
//	c.dbSome_app_id := flag.String("some_app_id", "testid", "db some_app_id")
//	c.dbSome_app_key := flag.String("some_app_key", "testkey", "db some_app_key")
//	return c
//}
