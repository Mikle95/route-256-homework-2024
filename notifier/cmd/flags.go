package main

import "flag"

type flags struct {
	topic             string
	bootstrapServer   string
	consumerGroupName string
}

func init() {
	flag.StringVar(&cliFlags.topic, "topic", "loms.order-events", "topic to produce")
	flag.StringVar(&cliFlags.bootstrapServer, "bootstrap-server", "localhost:9092", "kafka broker host and port")
	flag.StringVar(&cliFlags.consumerGroupName, "cg-name", "loms.order-events-consumer-group", "topic to produce")

	flag.Parse()
}
