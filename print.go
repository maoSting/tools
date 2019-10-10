package tools

import (
	"fmt"
	"strings"
)

func DebugPrint(tag string, args ...interface{}) {
	fmt.Println(tag)
	fmt.Println("---------------- start ----------------")
	for k, v := range args {
		fmt.Printf("%d: ", k)
		argType := fmt.Sprintf("%T", v)
		fmt.Print(argType, " ")
		if strings.ContainsAny(argType, "*") {
			fmt.Printf("%+v", v)
		} else {
			fmt.Printf("%v", v)
		}
		fmt.Println("")
	}
	fmt.Println("---------------- end ----------------")
}
