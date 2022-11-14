package utils

import (
	"math/rand"
	"time"
)

func RandSleep(duration int) {

	rand.Seed(time.Now().UnixNano())
	randSeconds := rand.Intn(duration)
	time.Sleep(time.Duration(randSeconds) * time.Second)

}

func RandSlice[T comparable](ids []T) int {
	rand.Seed(time.Now().UnixNano())
	randId := rand.Intn(len(ids) - 1)

	return randId

}
