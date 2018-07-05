package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// echo c29tZS1iYXNlLTY0 | b64 -d
// echo YzI5dFpTMWlZWE5sTFRZMA== | b64 -d -n 2
// echo some-base64 | b64
// b64 -d -in in.txt -out out.txt
// b64 -in in.txt -out out.txt

var (
	n   = flag.Int("n", 1, "number of iterations to be performed")
	v   = flag.Bool("v", false, "verbose, prints iterations")
	d   = flag.Bool("d", false, "decode")
	in  = flag.String("in", "", "input file")
	out = flag.String("out", "", "output file")
)

func main() {
	flag.Parse()

	if *in != "" {
		//open file
	}
	if *out != "" {
		//open file
	}

	if *d {
		reader := decode(os.Stdin, *n) // os.Stdin || or opened file
		all, _ := ioutil.ReadAll(reader)
		fmt.Println(string(all))
	} else {
		// encode()
	}
}

func decode(r io.Reader, n int) io.Reader {
	if n == 0 {
		return r
	}
	if *v {
		r = progress(r)
	}
	return decode(base64.NewDecoder(base64.StdEncoding, r), n-1)
}

func progress(r io.Reader) io.Reader {
	dst := bytes.NewBuffer([]byte{})
	io.Copy(dst, r)
	if s := strings.TrimSpace(string(dst.Bytes())); s != "" {
		fmt.Println(s)
	}
	return dst
}
