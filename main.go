package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int8 = -128
	fmt.Println("int8: ", a, "size of:", unsafe.Sizeof(a), "byte")
	var b int16 = 32767
	fmt.Println("int16: ", b, "size of:", unsafe.Sizeof(b), "byte")
	var c int64 = 999999999999999999
	fmt.Println("int64: ", c, "size of:", unsafe.Sizeof(c), "byte")
	
	var d float32 = 3.14
	fmt.Println("float32: ", d, "size of:", unsafe.Sizeof(d), "byte")

	var str string = "Hello worldefefefefsdvd wefefefefasdfasfsa "
	fmt.Println("string: ", str, "size of:", unsafe.Sizeof(str), "byte")
}
