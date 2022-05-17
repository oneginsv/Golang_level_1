package main

import (
	"flag"
	"fmt"
	"parser"
)

func main() {
	//	dbPort := flag.Int("port", 8080, "port for db")
	//	var dbConnectionURL string
	//	flag.StringVar(&dbConnectionURL, "db_url", "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable", "connection")
	//	dbJaeger_url := flag.String("jaeger_url", "http://jaeger:16686", "db jaeger_url")
	//	dbSentry_url := flag.String("sentry_url", "http://sentry:9000", "db sentry_url")
	//	dbKafka_broker := flag.String("kafka_url", "kafka:9092", "db kafka_url")
	//	dbSome_app_id := flag.String("some_app_id", "testid", "db some_app_id")
	//	dbSome_app_key := flag.String("some_app_key", "testkey", "db some_app_key")
	//	flag.Parse()

	config := parser.MyParser()

	//	fmt.Println(*dbPort)
	//	fmt.Println(dbConnectionURL)
	//	fmt.Println(*dbJaeger_url)
	//	fmt.Println(*dbSentry_url)
	//	fmt.Println(*dbKafka_broker)
	//	fmt.Println(*dbSome_app_id)
	//	fmt.Println(*dbSome_app_key)

	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("value of %s: %s\n", f.Name, f.Value.String())
	})
}
