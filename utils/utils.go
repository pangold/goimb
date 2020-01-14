package utils

import "time"

func GenerateMessageId() int64 {
	return time.Now().UnixNano()
}