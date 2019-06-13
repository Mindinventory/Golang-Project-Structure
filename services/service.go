package services

import (
	"crypto/rand"
	"fmt"
	"strings"
)

//RandomString returns the string in uppser-case
func RandomString(str string) string {
	b := make([]byte, 3)
	rand.Read(b)
	id := fmt.Sprintf(str+"%x", b)
	return strings.ToUpper(id)
}
