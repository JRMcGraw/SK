package main

import (
	"fmt"
)

type lots interface {
	f1()
	f2()
}

var j lots

func main() {


	fmt.Println(b())
	fmt.Println(c())
}

func b() lots {
	return &lots{ 
		f1:func(){
			fmt.Println(1)
		},
	}
}

func c() lots {
	return &lots{ 
		f2:func(){
			fmt.Println(1)
		},
	}
}

