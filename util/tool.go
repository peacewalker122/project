package util

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var (
	AlphaNumCheck = regexp.MustCompile(`^[a-zA-Z0-9_\s]+$`).MatchString
	AlphaCheck    = regexp.MustCompile(`^[a-zA-Z_\s]+$`).MatchString
	NumCheckByte  = regexp.MustCompile(`^[0-9]+$`).Match
	NumCheck      = regexp.MustCompile(`^[0-9]+$`).MatchString
	StringsCheck  = regexp.MustCompile(`^[a-zA-Z0-9_\s'"?!,.&%$@-]+$`).MatchString
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

type sqlInterface interface {
	sql.NullString
	sql.NullInt64
}

func InputSqlString(input string) (sql sql.NullString) {
	if input != "" {
		sql.Valid = true
		sql.String = input
		return sql
	}
	return sql
}

func RandomFileName(fileName string) (string, error) {

	ok := strings.ContainsRune(fileName, '.')
	if !ok {
		return "", fmt.Errorf("file is not valid")
	}

	_, filetype, _ := strings.Cut(fileName, ".")

	hasher := md5.New()
	hasher.Write([]byte(fileName))
	md5Hash := hex.EncodeToString(hasher.Sum(nil))

	unixTimestamp := time.Now().Unix()
	last4Digits := unixTimestamp % 10000
	md5Hash = md5Hash[:8]
	unique4string := Randomstring(4)

	filename := fmt.Sprintf("%s_%04d%s.%s", md5Hash, last4Digits, unique4string, filetype)
	return filename, nil
}

func RandomType() string {
	c := []string{"like", "retweet", "comment", "qoute-retweet"}
	l := len(c)
	return c[rand.Intn(l)]
}

func RandomUUID() string {
	return fmt.Sprintf("%v-%v-%v-%v", Randomstring(8), Randomstring(4), Randomstring(4), Randomstring(12))
}
