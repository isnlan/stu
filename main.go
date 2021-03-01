package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync/atomic"
	"time"
)

type User struct {
	name string
}

type C struct {
	s atomic.Value
}

func MakeMongoIdFromString(str string) string {
	bytes := sha256.Sum256([]byte(str))
	var dst [12]byte
	copy(dst[:], bytes[:])
	return hex.EncodeToString(dst[:])
}

func main() {
	t := time.Now().Unix()
	unix := time.Unix(t, 0)
	fmt.Println(unix)
}
