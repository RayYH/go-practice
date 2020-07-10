package main

import "fmt"

func ExampleDisplayAllKeywords() {
	validKeywords := []string{
		"break",
		"default",
		"func",
		"interface",
		"select",
		"case",
		"defer",
		"go",
		"map",
		"struct",
		"chan",
		"else",
		"goto",
		"package",
		"switch",
		"const",
		"fallthrough",
		"if",
		"range",
		"type",
		"continue",
		"for",
		"import",
		"return",
		"var",
	}

	for _, keyword := range validKeywords {
		fmt.Printf("%s\n", keyword)
	}

	// Output:
	// break
	// default
	// func
	// interface
	// select
	// case
	// defer
	// go
	// map
	// struct
	// chan
	// else
	// goto
	// package
	// switch
	// const
	// fallthrough
	// if
	// range
	// type
	// continue
	// for
	// import
	// return
	// var
}
