package internal

import (
	"math/rand"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var existedIds []string

func GenId(n int) string {
	for {
		rand.Seed(time.Now().UnixNano())
		b := make([]byte, n)
		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		if !isIdExisted(string(b)) {
			existedIds = append(existedIds, string(b))
			return string(b)
		}
	}
}

func isIdExisted(id string) bool {
	for _, val := range existedIds {
		if val == id {
			return true
		}
	}
	return false
}
