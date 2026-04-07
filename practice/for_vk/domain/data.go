package domain

import (
	"time"
)

func GetData() string {
	time.Sleep(10 * time.Millisecond)
	return "data"
}
