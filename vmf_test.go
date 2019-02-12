package vmf

import (
	"fmt"
	"os"
	"testing"
)

func TestVmf(T *testing.T) {
	T.Skip()
	file, _ := os.Open("maps/ze_candy_star_a1.vmf")

	reader := NewReader(file)
	f, _ := reader.Read()

	fmt.Println(f.Entities.GetChildrenByKey("entity")[147].GetProperty("classname"))
	fmt.Println(f.Entities.GetChildrenByKey("entity")[147].GetProperty("id"))
	fmt.Println(f.Entities.GetChildrenByKey("entity")[147].GetChildrenByKey("solid")[1])
}
