package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct {
}

func main() {
	resp, err := http.Get("http://www.google.com/")

	if err != nil {
		panic(err)
	}

	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes", len(bs))
	return len(bs), nil
}
