package vmf

import (
	"testing"
	"os"
)

func TestVmf(T *testing.T) {
	file,_ := os.Open("maps/ze_candy_star_a1.vmf")

	reader := NewReader(file)
	reader.Read()
}