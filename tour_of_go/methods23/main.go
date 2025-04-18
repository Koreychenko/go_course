package main

import (
	"io"
	"os"
	"strings"
)

const (
	rot13offset = 13
	enABCLength = 26
)

type rot13Reader struct {
	r io.Reader
}

func (rr *rot13Reader) Read(p []byte) (int, error) {
	n, err := rr.r.Read(p)

	for i := range p {
		switch {
		// We should process uppercase and lowercase letters separately
		case 'A' <= p[i] && p[i] <= 'Z':
			p[i] = 'A' + (p[i]-'A'+rot13offset)%enABCLength
		case 'a' <= p[i] && p[i] <= 'z':
			p[i] = 'a' + (p[i]-'a'+rot13offset)%enABCLength
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
