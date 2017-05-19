package main

import (
	"reflect"
	"fmt"
)

func showMeTheType(variable interface{}) string {
	return reflect.TypeOf(variable).String()
}

func main() {
	fmt.Println(showMeTheType(float64(7)));
}
