package go_files

import (
	"encoding/binary"
	"strings"
)

func Perform() {
	controlBits := make([]byte, 2)
	println(controlBits)
	binary.BigEndian.PutUint16(controlBits, uint16(0))
	println(controlBits)

	str := "abc::def::ghi"
	ss := strings.SplitN(str, "::", 2)
	println(str)
	println(ss[0])
	println(ss[1])
	println(len(ss))
	ss1 := strings.SplitAfterN(str, "::", 1)
	println(ss1[0])

}
