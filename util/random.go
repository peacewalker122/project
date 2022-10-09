package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Randomstring(n int) string {
	var sb strings.Builder
	l := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(l)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func Randomusername() string {
	return Randomstring(6)
}

func Randomemail() string {
	return fmt.Sprintf("%v@test.com", Randomstring(7))
}

func Randombyte() []byte{
	return []byte(Randomstring(7))
}
