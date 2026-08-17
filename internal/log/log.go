package log

import (
	"alfa_agent/internal/constants"
	"log"
)

func Info(message string) {
	log.Println(constants.Reset + message)
}

func Success(message string) {
	log.Println(constants.Green + message + constants.Reset)
}

func Warning(message string) {
	log.Println(constants.Reset + message + constants.Reset)
}

func Danger(message string) {
	log.Println(constants.Red + message + constants.Reset)
}

func Error(message error) {
	log.Println(constants.Red + message.Error() + constants.Reset)
}

func ErrorSpesific(format string, args ...any) {
	log.Printf(constants.Red+format, args...)
}