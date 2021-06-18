package main

import "fmt"

func main() {
	anInt := 1
	aFloat := 2.3
	aString := "Hello World"
	dict := map[string]interface{}{
		"name": "Ray",
		"age":  24,
	}
	person := struct {
		name string
		age  int
	}{
		name: "Ray",
		age:  24,
	}
	fmt.Println(anInt, aFloat, aString, dict, person)
}
