package main

import (
	"bytes"
	"fmt"
)

var str = []byte("abcdef")

func seekorigin() {
	fmt.Println("Origin:")
	b := bytes.NewReader(str)
	fmt.Println(b.Seek(1, 0)) //0 means relative to the origin of the file
	c, _ := b.ReadByte()
	fmt.Println(string(c))
}

func seekcurrent() {
	fmt.Println("Current:")
	b := bytes.NewReader(str)
	b.ReadByte()              // read next byte
	b.ReadByte()              // read another byte
	fmt.Println(b.Seek(1, 1)) // 1 means relative to the current offset
	c, _ := b.ReadByte()
	fmt.Println(string(c))
}

func seekend() {
	fmt.Println("End:")
	b := bytes.NewReader(str)
	fmt.Println(b.Seek(-1, 2)) //2 means relative to the end
	c, _ := b.ReadByte()
	fmt.Println(string(c))
}

func main() {
	seekorigin()
	seekcurrent()
	seekend()
}
