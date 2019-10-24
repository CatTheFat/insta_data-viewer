package main

import (
	"fmt"
	"runtime"
)

//ErrHandle : prints an error if there is one
func ErrHandle(err error) {
	if err != nil {
		fmt.Println(err)

		_, _, line, ok := runtime.Caller(1)
		if ok {
			fmt.Printf("Error check called from #%d\n", line)
		}
	}
}
