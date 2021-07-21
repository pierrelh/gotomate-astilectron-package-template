/*
The template.go is used for building the dialogs that are shown to the user when they click on the instruction

All you have to do is to list your templates with the needed fields like this:

var ATemplate = &template.InstructionTemplate{
	...
}

var BTemplate = &template.InstructionTemplate{
	...
}
...

For better lisibility, list them in alphabetical order & always set the name starting by an uppercase

Their are bunch of differents type of input that you can use for your template
*/

package yourpackagename

import (
	"gotomate-astilectron/fiber/template"
)

// A way to set some select's options
func MyOptions() []template.Option {
	return []template.Option{
		{Name: "name1", Value: "value1"},
		{Name: "name2", Value: "value2"},
		{Name: "name3", Value: "value3"},
	}
}

// template.Select{
// 	Bind:         "FooString",
// 	BindVariable: "FooStringVarName",
// 	Options:      MyOptions(), // Then call your options like this
// }

// YourFunctionTemplate Dialog's Template of your function
var YourFunctionTemplate = &template.InstructionTemplate{
	template.Field{ // Creating a new Field in the Template for Foo String
		Label: template.Label{ // Creating the label of the field Foo String
			Text: "Foo String:", // Setting the text of the label
		},
		Input: template.TextInput{ // Creating the field foo wich is a string
			Bind:         "FooString",        // Binding the raw value according to the databinder
			BindVariable: "FooStringVarName", // Binding the variable name value according to the databinder
			// Value: "",					  // Use this to set a default value
			// Disabled: true, 				  // Use this to disable the input
			// Prefix: "",					  // Use this to set a prefix on the value
			// Suffix: "",					  // Use this to set a suffix on the value
			// MaxLength: 2,				  // Use this to set a max length to your TextInput
			// MinLength: 1,,				  // Use this to set a min length to your TextInput
		},
		VariableToggler: template.VariableToggler{ // Creating the Variable toggler (if foo string can be a var)
			Bind: "FooStringIsVar", // Binding the value according to the databinder
		},
	}.Build(), // Always finish a template.Field{...} by .Build()
	template.Field{ // Creating a new Field in the Template for Foo Float
		Label: template.Label{
			Text: "Foo Float:",
		},
		Input: template.NumberInput{
			Bind:         "FooFloat",        // Binding the raw value according to the databinder
			BindVariable: "FooFloatVarName", // Binding the variable name value according to the databinder
			Decimals:     5,                 // Set the number of decimals of your float
			// Value: "", 					 // Use this to set a default value
			// Prefix: "", 					 // Use this to set a prefix on the value
			// Suffix: "", 					 // Use this to set a suffix on the value
			// Disabled: false, 			 // Use this to disable the input
			// MinValue: 1, 				 // Use this to set a min value to your NumberInput
			// MaxValue: 2,					 // Use this to set a max value to your NumberInput
		},
		VariableToggler: template.VariableToggler{
			Bind: "FooFloatIsVar", // Binding the value according to the databinder
		},
	}.Build(),
	template.Field{ // Creating a new Field in the Template for Foo Int
		Label: template.Label{
			Text: "Foo Int:",
		},
		Input: template.NumberInput{
			Bind:         "FooInt",        // Binding the raw value according to the databinder
			BindVariable: "FooIntVarName", // Binding the variable name value according to the databinder
			Decimals:     0,               // Set the number of decimals to 0 to get an int
			// Value: "", // Use this to set a default value
			// Prefix: "", // Use this to set a prefix on the value
			// Suffix: "", // Use this to set a suffix on the value
			// Disabled: false, // Use this to disable the input
			// MinValue: 1, // Use this to set a min value to your NumberInput
			// MaxValue: 2, // Use this to set a max value to your NumberInput
		},
		VariableToggler: template.VariableToggler{
			Bind: "FooIntIsVar", // Binding the value according to the databinder
		},
	}.Build(),
	template.Field{ // Creating a new Field in the Template for Foo Bool
		Label: template.Label{
			Text: "Foo Bool:",
		},
		Input: template.CheckboxInput{
			Bind:         "FooBool",        // Binding the raw value according to the databinder
			BindVariable: "FooBoolVarName", // Binding the variable name value according to the databinder
			// Value: "", // Use this to set a default value
			// Disabled: false, // Use this to disable the input
		},
		VariableToggler: template.VariableToggler{
			Bind: "FooBoolIsVar", // Binding the value according to the databinder
		},
	}.Build(),
	template.Field{ // Creating a new Field in the Template for Foo Only Var Name
		Label: template.Label{
			Text: "Foo Only Var Name:",
		},
		Input: template.TextInput{
			BindVariable: "FooOnlyVarName", // Binding the variable name value according to the databinder
			// Value: "", // Use this to set a default value
			// Disabled: false, // Use this to disable the input
		},
		VariableToggler: template.VariableToggler{
			Checked:  true, // Set the variable checkbox to true by default for setting a varible
			Disabled: true, // Disable the checkbox to understand that the field must be a variable
		},
	}.Build(),
	template.Field{ // Creating a new Field in the Template for Ouput (giving a name of a value created by your functions)
		Label: template.Label{
			Text: "Output:",
		},
		Input: template.TextInput{
			BindVariable: "Output", // Binding the variable name value according to the databinder
		},
	}.Build(),
}
