/*
The functions.go file is used for declaring your functions & fill them with what you
want the automate to do.
Please keep the functions name's with the first letter as an uppercase

For better lisibility, & user experience list your functions in alphabetical order because Gotomate will get them
by the order that you set.
*/

package yourpackagename

import (
	"gotomate-astilectron/fiber/variable"
	"gotomate-astilectron/log"
)

// Define a function
// instructionData -> Get the data to the associated instruction (values, var, etc.)
// finished -> Channel of bool to tell the fiber when the function is finished
// The function return an int to tell the fiber the next instruction id to use, return -1 for default
func YourFunction(instructionData interface{}, finished chan bool) int {
	log.FiberInfo("Doing Something")

	// To get a value from the fiber
	// Send to the function the instructions datas & the differents databinder fields of the value
	foo, err := variable.Keys{VarName: "FooVarName", Name: "Foo", IsVarName: "FooIsVar"}.GetValue(instructionData)
	if err != nil {
		// Return if the value isn't found
		finished <- true
		return -1
	}

	// To set / update a value in the fiber do so:
	// Always set the instructionData first the the key to set / update, then the new value
	variable.SetVariable(instructionData, "SomeKey", foo)

	// Do what you want here

	// Always finish your functions by this:
	finished <- true // Send to the channel that the function is finished
	return -1        // Return the next instruction id, return -1 for default
}
