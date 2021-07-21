/*
The processing.go file is used to list all your functions in functions.go & process them when called by the fiber.

All you have to do is to list your functions like this:
case "A":
	go func() {
		nextID = A(instructionData, finished)
	}()
	<-finished
case "B":
	go func() {
		nextID = B(instructionData, finished)
	}()
	<-finished
...

For better lisibility, list them in alphabetical order
*/

package yourpackagename

// Processing the functions.go functions
func Processing(funcName string, instructionData interface{}, finished chan bool) int {
	nextID := -1
	switch funcName {
	case "YourFunction":
		go func() {
			nextID = YourFunction(instructionData, finished)
		}()
		<-finished
	default:
		return -2
	}
	return nextID
}
