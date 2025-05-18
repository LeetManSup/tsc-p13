package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	nc.QueueSubscribe("jobs.create", "workers", func(m *nats.Msg) {
		fmt.Printf("Обработано: %s\n", string(m.Data))
	})

	select {} // Блокируем завершение
}
