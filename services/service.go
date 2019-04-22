package services

import (
	"crypto/rand"
	"fmt"
	"strings"
)

//Common functions
func RandomString(str string) string {
	b := make([]byte, 3)
	rand.Read(b)
	id := fmt.Sprintf(str+"%x", b)
	return strings.ToUpper(id)
}
