package common

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

type BaseCtrl struct {
	// LogCtrl
}

func (e *BaseCtrl) Md5(s string) string {
	hash := md5.New()
	buf := []byte(s)
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
func (e *BaseCtrl) Sha1(s string) string {
	hash := sha1.New()
	buf := []byte(s)
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

//Sha1 Plus Md5
func (e *BaseCtrl) Sha1PlusMd5(s string) string {
	return e.Sha1(e.Md5(s))
}
