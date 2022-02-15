// EXAPLE OF HTTP PACKAGE IN GO LENGUAGE BY DANI95RICO

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	fmt.Println("Se han escrito esta cantidad de bytes: ", len(bs))
	return len(bs), nil
}
