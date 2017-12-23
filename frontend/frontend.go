package frontend

import (
	"github.com/robertkrimen/otto"
)

// ExampleExecute example execute ES5.1
func ExampleExecute() int64 {
	vm := otto.New()
	vm.Run(`
		try { eval('"use strict"; class foo {}'); } catch (e) { console.log(e); }

    abc = 2 + 2;
	console.log("The value of abc is " + abc); // 4`)
	val, _ := vm.Get("abc")

	va, _ := val.ToInteger()

	return va
}
