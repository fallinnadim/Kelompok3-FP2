package helper

import "log"

func ErrorFatal(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}
