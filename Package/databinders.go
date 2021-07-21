/*
The databinders.go file is used for declaring the datas needed to link the dialogs to the function

All you have to do is to list your databinder with the needed paramaters like this:

type ADatabinder struct {
	paramA string
	...
}

type BDatabinder struct {
	paramB string
	...
}
...

For better lisibility, list them in alphabetical order & always set the name starting by an uppercase
*/

package yourpackagename

// YourFunctionDatabinder Define the parameters needed for your function & that will be filled in the template
type YourFunctionDatabinder struct {
	FooString        string
	FooStringVarName string
	FooStringIsVar   bool
	FooFloat         float64
	FooFloatVarName  string
	FooFloatIsVar    bool
	FooInt           int
	FooIntVarName    string
	FooIntIsVar      bool
	FooBool          bool
	FooBoolVarName   string
	FooBoolIsVar     bool
	FooOnlyVarName   string
	Output           string
}
