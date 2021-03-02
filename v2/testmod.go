package testmod

import (
	"fmt"
	_ "github.com/valyala/fasthttp"
)
//Важная деталь: каталог пакета следует разместить за пределами вашего $GOPATH,
//потому что, внутри него по умолчанию отключена поддержка модулей.
// Hi returns a friendly greeting
func Hi(name string) string {
	switch name {
	case "Murlitas":
		return fmt.Sprintf("Hi, %s", name)
	default:
		return fmt.Sprintf("Hello, %s", name)
	}

}
