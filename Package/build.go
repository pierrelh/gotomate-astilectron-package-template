/*
The build.go file is used to get the right databinder & template when a new instruction is builded,
& also used to decode a saved fiber by getting the databinder to use.

All you have to do is to list your functions with the associated databinder & template like this:
case "A":
	return new(ADatabinder), ATemplate
case "B":
	return new(BDatabinder), BTemplate
...

For better lisibility, list them in alphabetical order
*/
package yourpackagename

import (
	"gotomate-astilectron/fiber/template"
)

// Build Return the right databinder & the right template for a yourpackage instruction
func Build(function string) (interface{}, *template.InstructionTemplate) {
	switch function {
	case "YourFunction":
		return new(YourFunctionDatabinder), YourFunctionTemplate
	}
	return nil, nil
}
