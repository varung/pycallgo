package main

// #include <stdlib.h>
import "C"
import (
	"encoding/json"
	"fmt"
	"log"
	"unsafe"
)

type Event struct {
	Operation string
	Value     float64
}

func Handle(e Event) string {
	return fmt.Sprintf("%f", e.Value*2)
}

// Enable python to call it

//export HandleRequest
func HandleRequest(datap *C.char) *C.char {
	data := C.GoString(datap)
	log.Println(data)
	var e Event
	err := json.Unmarshal([]byte(data), &e)
	log.Println(err)
	if err != nil {
		return C.CString("Could not parse json: " + data)
	}
	res := Handle(e)
	return C.CString(res)
}

//export Free
func Free(datap *C.char) {
	C.free(unsafe.Pointer(datap))
}

func main() {
}
