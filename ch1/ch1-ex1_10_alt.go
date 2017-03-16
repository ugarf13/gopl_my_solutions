// Fetchall fetches URLs in parallel and reports their time and sizes.

package main

import (
	"fmt"
	//"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	file, err := os.OpenFile("fetch.txt", os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0744)
	if err != nil {
		fmt.Fprintf(os.Stderr, "OpenFile: %v\n", err)
	}
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch, file) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, file *os.File) {
	start := time.Now()
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close() // don't leak resources
	//nbytes, err := io.Copy(file, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	nbytes, err := file.Write(b)
	if err != nil {
		ch <- fmt.Sprintf("while writing %s: %v", url, err)
		return
	}
	err = file.Close()
	if err != nil {
		ch <- fmt.Sprintf("close: %v\n", err)
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
