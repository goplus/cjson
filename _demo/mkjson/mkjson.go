package main

import (
	"github.com/goplus/cjson"
	"github.com/goplus/llgo/c"
)

func main() {
	mod := cjson.Object()
	mod.SetItem(c.Str("name"), cjson.String(c.Str("math")))
	c.Printf(c.Str("%s\n"), mod.CStr())
	mod.Delete()
}
