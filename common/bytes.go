package common

import (
	"encoding/hex"
)

// FromHex returns the bytes represented by the hexadecimal string s.
func FromHex(s string) []byte {
	if len(s) > 1 {
		if s[0:2] == "0x" || s[0:2] == "0X" {
			s = s[2:]
		}
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}
	return Hex2Bytes(s)
}

// Hex2Bytes returns the bytes represented by the hexadecimal string str.
func Hex2Bytes(str string) []byte {
	// hex.DecodeString  f func(s string) ([]byte, error)
	h, _ := hex.DecodeString(str)
	return h
}
