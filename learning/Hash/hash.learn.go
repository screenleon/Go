package main

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
)

func main() {
	h := crc32.NewIEEE()
	h.Write([]byte("test"))
	v := h.Sum32()
	fmt.Println(v)

	h2 := sha1.New()
	h2.Write([]byte("test"))
	bs := h2.Sum([]byte{})
	fmt.Println(bs)
}
