package vmf

import (
	"testing"
	"os"
	"fmt"
)

func TestVmf(T *testing.T) {
	file,_ := os.Open("maps/ze_candy_star_a1.vmf")

	reader := NewReader(file)
	f,_ := reader.Read()

	fmt.Println(f.Entities.GetChildrenOfType("entity")[147].GetProperty("classname"))
	fmt.Println(f.Entities.GetChildrenOfType("entity")[147].GetProperty("id"))
	fmt.Println(f.Entities.GetChildrenOfType("entity")[147].GetChildrenOfType("solid")[1])
}