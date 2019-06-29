package service

// Parameter describes task input parameters, output parameters of a task
// output and input parameters of an event.
type Parameter struct {
	// Key is the key of parameter.
	Key string `hash:"name:1" validate:"printascii"`

	// Name is the name of parameter.
	Name string `hash:"name:2" validate:"printascii"`

	// Description is the description of parameter.
	Description string `hash:"name:3" validate:"printascii"`

	// Type is the data type of parameter.
	Type string `hash:"name:4" validate:"required,printascii,oneof=String Number Boolean Object Any"`

	// Optional indicates if parameter is optional.
	Optional bool `hash:"name:5"`

	// Repeated is to have an array of this parameter
	Repeated bool `hash:"name:6"`

	// Definition of the structure of the object when the type is object
	Object []*Parameter `hash:"name:7" validate:"unique,dive,required"`
}
