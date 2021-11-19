package util

import (
	"github.com/astaxie/beego/validation"
	"log"
)

func PrintLog(valid *validation.Validation) {
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Printf("Validation err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
}
