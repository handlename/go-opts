package main

import (
	"fmt"

	"github.com/handlename/go-opts"
)

type myOpts struct {
	Name    string  `flag:"name" default:"jhon" description:"your name"`
	Age     int     `flag:"age" default:"25" description:"your age"`
	Weight  float64 `flag:"weight" default:60.0 description:"your weight"`
	Working bool    `flag:"working" default:false description:"do you working?"`
}

func main() {
	o := myOpts{}
	opts.Parse(&o)
	fmt.Println(o)
}
