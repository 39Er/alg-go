package crypto

import (
	"bytes"
	"errors"
)

var (
	PKCS7 = &pkcs7{}
	ErrPaddingSize = errors.New("padding size error")
)

type pkcs7 struct{}

func (*pkcs7) padding(src []byte, blockSize int)(result []byte) {
	srcLen := len(src)
	padding := blockSize - (srcLen % blockSize)
	pads:=bytes.Repeat([]byte{byte(padding)}, padding)
	result = append(src,pads...)
	return
}

func (*pkcs7)unpadding(src []byte, blockSize int)( []byte, error){
	srcLen:=len(src)
	padding:=int(src[srcLen-1])
	if padding>srcLen || padding >blockSize {
		return nil,ErrPaddingSize
	}
	return src[:srcLen-padding], nil
}
