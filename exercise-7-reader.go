package main

import "golang.org/x/tour/reader"
import "fmt"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (MyReader) Read(b []byte) (n int, e error) {
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
	
	//extra
	r := MyReader{}
	b := make([]byte, 8)
	n, err := r.Read(b)
	fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
	fmt.Printf("b[:n] = %q\n", b[:n])
}
