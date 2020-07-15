package main

import (
	"bytes"
	"fmt"
)

/*
	https://juejin.im/post/5ef3108af265da22a77c3a5c
*/

func main() {
	buf1 := bytes.NewBufferString("hello")
	buf2 := bytes.NewBuffer([]byte("hello"))
	buf3 := bytes.NewBuffer([]byte{'h','e','l','l','o'})
	fmt.Printf("%v, %v, %v\n", buf1, buf2, buf3)
	fmt.Printf("%v, %v, %v\n", buf1.Bytes(), buf2.Bytes(), buf3.Bytes())

	buf4 := bytes.NewBufferString("")
	buf5 := bytes.NewBuffer([]byte{})
	fmt.Println(buf4.Bytes(), buf5.Bytes())
}
