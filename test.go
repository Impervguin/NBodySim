package main

import (
	"fmt"
	"sync"
	"unsafe"
)

func main() {
	m := sync.Mutex{}
	fmt.Println(unsafe.Sizeof(m))
}
