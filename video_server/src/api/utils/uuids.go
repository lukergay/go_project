package utils

import (
	"crypto/rand"
	"fmt"
	"io"
)

//生成唯一标识符
func NewUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	//b变量的id
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[8]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])

}
