package _slices

import "bytes"

func Append(slice, data []byte) []byte {
	//return append(slice, data...)
	buf := bytes.NewBuffer(slice)
	buf.Write(data)

	return buf.Bytes()
}
