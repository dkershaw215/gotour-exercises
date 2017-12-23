package main

import (
	"io"
	"os"
	"strings"
)

var key = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
var lookup = []byte("NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm")

var rot13Map = make(map[byte]byte)

type rot13Reader struct {
	r io.Reader
}

func buildRot13Map() {
	for i:=0; i < len(key); i++ {
		rot13Map[key[i]] = lookup[i]
	}
}

func rot13(b byte) byte {
	if c, ok := rot13Map[b]; ok {
		return c
	}
	return b
}

func (r13 rot13Reader) Read(b []byte) (n int, e error) {
	n, e = r13.r.Read(b)
	for i := 0; i < n; i++ {
        b[i] = rot13(b[i])
    }
    return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	buildRot13Map()
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
