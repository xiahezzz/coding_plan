package user

import (
	"bytes"
)

func ConcatString(a string, b string) string {
	/*
		利用buffer来操作/增加字符串效率较高
		如果能预估字符串长度，还可以用
		buffer.Grow()
		来设置capacity
	*/

	var buffer bytes.Buffer

	buffer.WriteString(a)
	buffer.WriteString(b)

	return buffer.String()
}
