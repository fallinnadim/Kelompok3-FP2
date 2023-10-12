package helper

import (
	"fmt"
	"time"
)

func CurrentDate() time.Time {
	layoutFormat := "2006-01-02"
	date, _ := time.Parse(layoutFormat, time.Now().String())
	fmt.Println(time.Now().Format("2006-01-02"))
	return date
}
