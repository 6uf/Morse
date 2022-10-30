package Morse

import (
	"strings"
)

func Encode(input string) (build string) {
	for _, char := range input {
		if val, ok := Morse[strings.ToLower(string(char))]; ok {
			build += " " + val + " "
		}
	}
	return build[1 : len(build)-1]
}

func Decode(M string) (Decode string) {
	for _, V := range strings.Split(M, " ") {
		if val, ok := ReverseMorse[V]; ok {
			Decode += val
		}
	}
	return
}
