[![GoDoc](https://godoc.org/github.com/Galaco/vmf?status.svg)](https://godoc.org/github.com/Galaco/vmf)
[![Go report card](https://goreportcard.com/badge/github.com/galaco/vmf)](https://goreportcard.com/badge/github.com/galaco/vmf)
[![Build Status](https://travis-ci.com/Galaco/vmf.svg?branch=master)](https://travis-ci.com/Galaco/vmf)

# vmf
Golang library for manipulating hammer Editor .vmf map files

This library constructs a simple node tree that you can query any structure(s) and any property(s) of.

### Usage
```golang
package main

import (
    "fmt"
    "github.com/galaco/vmf"
)

func main() {
	file,_ := os.Open("de_dust2.vmf")

	reader := vmf.NewReader(file)
	f,_ := reader.Read()

	fmt.Println(f.Entities.GetChildrenByKey("entity")[0].GetProperty("classname"))
	fmt.Println(f.Entities.GetChildrenByKey("entity")[0].GetChildrenByKey("solid"))
}
```