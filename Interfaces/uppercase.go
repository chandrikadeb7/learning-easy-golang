package main

import (
	"fmt"
	"io"
	"os"
)

//IO writer field
type Capper struct {
	wtr io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {
	//difference of char is a byte in Go
	diff := byte('a' - 'A')

	out := make([]byte, len(p)) //buffer
	for i, c := range p {
		if c >= 'a' && c <= 'z' { //if lowercase letter
			c -= diff
		}
		out[i] = c //output upper case letter
	}
	return c.wtr.Write(out)
}

func main() {
	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "Hello, world!")

}
