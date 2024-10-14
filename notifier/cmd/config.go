package main

import "gitlab.ozon.dev/1mikle1/homework/notifier/internal/infra/kafka"

type config struct {
	KafkaConfig kafka.Config
}

func newConfig(f flags) config {
	return config{
		KafkaConfig: kafka.Config{
			Brokers: []string{
				f.bootstrapServer,
			},
		},
	}
}
