package kzfile

import (
	"fmt"
	"time"
)

func GenerateRandomTime() string {
	now := time.Now()
	return fmt.Sprintf("%v%v%v%v%v%v", now.Year(), int(now.Month()), now.Day(), now.Hour(), now.Minute(), now.Second())
}
