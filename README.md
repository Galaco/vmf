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

	fmt.Println(f.Entities[0].GetProperty("classname"))
	fmt.Println(f.Entities[0].GetChildrenOfType("solid"))
}
```