package main

import (
	"fmt"
)

type Optioner interface {
	Name() string
	IsValid() bool
}

// structs with tags field.
// tags have no functional purpose but unlike comments
// they are accessible using Go’s reflection support
type OptionCommon struct {
	ShortName string "short option name"
	LongName  string "long option name"
}

// The name methods id common to all object
// that embed OptionCommon
func (option OptionCommon) Name() string {
	if option.LongName == "" {
		return option.ShortName
	}
	return option.LongName
}

// -------------------
// Int Option
type IntOption struct {
	OptionCommon
	Value, Min, Max int
}

//
// let IntOption satisfy the Optioner interface
// implementing IsValid method
func (option IntOption) IsValid() bool {
	return option.Min <= option.Value && option.Value <= option.Max
}

// -------------------
// String Option

//
// define a new type string option
type StringOption struct {
	OptionCommon
	Value, Min, Max string
}

//
// let StringOption satisfy the Optioner interface
// implementing the IsValid method
func (option StringOption) IsValid() bool {
	return option.Min <= option.Value && option.Value <= option.Max
}

func main() {

	// object createion
	i := IntOption{OptionCommon{"short_int", "long_int"}, 10, 0, 100}
	s := StringOption{OptionCommon{"short_string", "long_string"}, "10", "0", "100"}

	fmt.Println("Hello embedding interface", i.Name(), i.Value)
	fmt.Println("Hello embedding interface", s.Name(), s.Value)

	o := []Optioner{i, s}
	for _, option := range o {

		fmt.Println("Hello embedding interface", option.Name(), option.IsValid())

		// to access the value we must cast the option
		// to the real value
		switch option := option.(type) { // shadow variable
		case IntOption:
			fmt.Println("Value Int", option.Value)
		case StringOption:
			fmt.Println("Value String", option.Value)
		}
	}
}
