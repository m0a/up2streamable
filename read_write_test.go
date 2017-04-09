package main

import (
	"fmt"
	"io"
	"testing"
)

type MyIO chan byte

func (io MyIO) listenChannel() {
	for out := range io {
		fmt.Println(
			fmt.Sprintf("out from channel => %d", out),
		)
	}
}

func (io MyIO) Write(p []byte) (n int, err error) {
	n = 0
	for _, b := range p {
		io <- b
		n++
	}
	return n, nil
}

func (io MyIO) Read(p []byte) (n int, err error) {

	return 0, nil
}

var _ io.Writer = make(MyIO)
var _ io.Reader = make(MyIO)

// func listenChannel(myCh chan int) {
// 	// myChからの送信を受け取る
// 	for out := range myCh {
// 		fmt.Println(
// 			fmt.Sprintf("out from channel => %d", out),
// 		)
// 	}
// }

func TestSampleTest(t *testing.T) {

	r, w := io.Pipe()

	go fmt.Fprintln(w, "Hello, World")
	var buf = make([]byte, 64)
	n, err := r.Read(buf)
	if err != nil {
		t.Errorf("read: %v", err)
	} else if n != 12 || string(buf[0:12]) != "hello, world" {
		t.Errorf("bad read: got %q", buf[0:n])
	}

	// myChan := make(MyIO)

	// // go listenChannel(myChan)
	// go myChan.listenChannel()

	// fmt.Fprintf(myChan, "1234")

	t.Error("Hogeee.")
}
