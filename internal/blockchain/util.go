package blockchain

import (
	"bytes"
	"encoding/binary"
	"log"
)

func IntToByte(o int64) []byte {
	buf := bytes.NewBuffer([]byte{})
	err := binary.Write(buf, binary.BigEndian, o)
	if err != nil {
		log.Fatalln(err)
	}
	return buf.Bytes()
}
