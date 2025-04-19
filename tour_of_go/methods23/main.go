package main

import (
	"fmt"
	"io"
	"strings"
)

const (
	rot13       = 13
	alphabetLen = 26
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(p []byte) (int, error) {
	n, err := r13.r.Read(p)

	for i := 0; i < n; i++ {
		switch {
		// We should process uppercase and lowercase letters separately
		case 'A' <= p[i] && p[i] <= 'Z':
			p[i] = 'A' + (p[i]-'A'+rot13)%alphabetLen
		case 'a' <= p[i] && p[i] <= 'z':
			p[i] = 'a' + (p[i]-'a'+rot13)%alphabetLen
		}
	}

	return n, err
}

func Rot13Transform(str string) (string, error) {
	s := strings.NewReader(str)
	r := rot13Reader{s}
	converted, err := io.ReadAll(r)

	return string(converted), err
}

func main() {
	converted, err := Rot13Transform("Lbh penpxrq gur pbqr!")

	if err != nil {
		panic(err)
	}

	fmt.Println(converted)
}
