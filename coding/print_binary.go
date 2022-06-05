package coding

import (
	"bytes"
	"fmt"
)

func PrintBinary(a int) {
	var buffer bytes.Buffer
	buffer.Grow(32)
	for i := 31; i > 0; i-- {
		if a&(1<<i) == 0 {
			buffer.WriteString("0")
		} else {
			buffer.WriteString("1")
		}

	}
	fmt.Printf("res: %v\n", buffer.String())
}
