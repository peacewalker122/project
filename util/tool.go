package util

import (
	"database/sql"
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

func Randombyte() []byte {
	return []byte(Randomstring(7))
}

func Randomint(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func Randomuint() []uint8 {
	return []uint8(Randomstring(90))
}

func InputSqlString(input string) (sql sql.NullString) {
	if input != "" {
		sql.Valid = true
		sql.String = input
		return sql
	}
	return sql
}

func RandomType() string {
	c := []string{"like", "retweet", "comment", "qoute-retweet"}
	l := len(c)
	return c[rand.Intn(l)]
}
