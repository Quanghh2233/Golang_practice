package main

import (
	"fmt"
)

func main() {
	WhatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Dont know type %T\n", t)
		}
	}
	WhatAmI(true)
	WhatAmI(1)
	WhatAmI("Pretty good")
}
