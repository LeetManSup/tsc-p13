package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("nats://nats:4222")
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	for i := 1; i <= 10; i++ {
		msg := fmt.Sprintf("Task %d", i)
		err := nc.Publish("jobs.create", []byte(msg))
		if err != nil {
			fmt.Println("Ошибка при отправке:", err)
		} else {
			fmt.Println("Отправлено:", msg)
		}
		time.Sleep(1 * time.Second)
	}
}
