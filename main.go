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
// b64 -d
// b64

var (
	n = flag.Int("n", 1, "number of iterations to be performed")
	v = flag.Bool("v", false, "verbose, prints iterations")
	d = flag.Bool("d", false, "decode")
)

func main() {
	flag.Parse()

	if *d {
		reader := decode(os.Stdin, *n)
		all, _ := ioutil.ReadAll(reader)
		fmt.Println(string(all))
	} else {
		out := encode(os.Stdout, *n)
		defer out.Close()

		clear, _ := ioutil.ReadAll(os.Stdin)
		out.Write(clear)
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

func encode(r io.WriteCloser, n int) io.WriteCloser {
	if n == 0 {
		return r
	}
	if *v {

	}
	return encode(base64.NewEncoder(base64.StdEncoding, r), n-1)
}
