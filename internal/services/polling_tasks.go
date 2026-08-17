package services

import (
	"alfa_agent/internal/log"
	"time"
)

var (
	is_running = false
)

func StartingPolling(interval int){
	is_running = true

	for is_running {
		tasks_polling()
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func tasks_polling() {
	log.Info("Polling tasks")
}