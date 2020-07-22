package main

// typedef void (*callbackFunc) (void * ctx,const char *json);
//
// static inline void
// bridge_callback_func(callbackFunc f,void * ctx,const char* json)
// {
//		f(ctx,json);
// }
//
import "C"
import (
	"time"
	"unsafe"
)

// Callback struc used for storing context and callback function
type Callback struct {
	fn  C.callbackFunc
	ctx unsafe.Pointer
}

// GetCallback create a function who call the cgo bridge and set parameter
func (c *Callback) GetCallback() func(string) {
	return func(json string) {
		C.bridge_callback_func(c.fn, c.ctx, C.CString(json))
	}
}

// cb don't do that in a real apps
var cb Callback

//export RegisterCallBack
func RegisterCallBack(handle unsafe.Pointer, ctx unsafe.Pointer) {
	cb = Callback{fn: C.callbackFunc(handle), ctx: ctx}
}

//export Start
// call a simple goroutine from C Code
// just for showing parallel execution
func Start() {
	go exec()
}

func exec() {
	time.Sleep(3 * time.Second)
	fn := cb.GetCallback()
	// Call the C Callback
	fn("cool")
}

// main mandatory but not use there
func main() {}
