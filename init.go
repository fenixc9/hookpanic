package hookpanic

import (
	"fmt"
	"git.aimap.io/hao.liu/hook"
	"runtime"
	_ "runtime"
	"runtime/debug"
	_ "unsafe"
)

//go:linkname gopanic runtime.gopanic
func gopanic(a interface{})

type handler func(a interface{}, stack []byte)

var defaultHandler handler

//go:noinline
func hookPanic(a interface{}) {
	defaultHandler(a,debug.Stack())
	origingopanic(a)
}

//go:noinline
func origingopanic(a interface{}) {
	fmt.Println("if this prints , its a bug")
	runtime.Goexit()
}

func SetPanicHandler(f handler) {
	defaultHandler = f
	hook := hook.Hook(gopanic, hookPanic, origingopanic)
	if hook != nil {
		fmt.Println("hook panic 失败:", hook.Error())
	}
}
