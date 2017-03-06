package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

const base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

var coder = base64.NewEncoding(base64Table)

func base64Encode(encode_byte []byte) []byte {
	return []byte(coder.EncodeToString(encode_byte))
}

func base64Decode(decode_byte []byte) ([]byte, error) {
	return coder.DecodeString(string(decode_byte))
}

func MD5(b []byte) string {

	h := md5.New()

	h.Write(b)

	x := h.Sum(nil)

	y := make([]byte, 32)

	hex.Encode(y, x)

	return string(y)

}

func main() {

	encode_str := "博客地址：http://blog.csdn.net/bojie5744"

	decode_byte := base64Encode([]byte(encode_str))

	fmt.Printf("%s\n", string(decode_byte))

	encode_byte, _ := base64Decode(decode_byte)

	fmt.Println(string(encode_byte))

	fmt.Println(MD5([]byte(encode_str)))

}
