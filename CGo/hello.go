package main

//void SayHello(const char* s);
import (
	"C"
)


func main(){
	println("hello cgo")
	// cstr := C.CString("Hello, World\n")
	// C.puts(cstr)
	C.SayHello(C.CString("hello, world\n"))
}
