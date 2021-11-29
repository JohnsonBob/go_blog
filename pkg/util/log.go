package util

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"go_blog/pkg/logging"
	"log"
)

func PrintLog(valid *validation.Validation) {
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Printf("Validation err.key: %s, err.message: %s", err.Key, err.Message)
			logging.Info(err.Key, err.Message)
		}
	}
}

func Printf(format string, v ...interface{}) {
	log.Printf(format, v)
	logging.Info(fmt.Printf(format, v))
}

func Println(v ...interface{}) {
	log.Println(v)
	logging.Info(v)
}

func Fatal(v ...interface{}) {
	logging.Info(v)
	log.Fatal(v)
}

func Fatalf(format string, v ...interface{}) {
	logging.Info(fmt.Sprintf(format, v...))
	log.Fatal(v)
}
