package db0102

import (
	"io"
	"encoding/binary"
	// "fmt"
)

type Entry struct {
	key []byte
	val []byte
}

func (ent *Entry) Encode() []byte {
	data := make([]byte, 4 + 4 + len(ent.key) + len(ent.val))
	binary.LittleEndian.PutUint32(data[0:4], uint32(len(ent.key)))
	binary.LittleEndian.PutUint32(data[4:8], uint32(len(ent.val)))
	copy(data[8:], ent.key)
	copy(data[8 + len(ent.key):], ent.val)
	return data
}

//old first implementation
// func (ent *Entry) Decode(r io.Reader) error {
// 	header := make([]byte, 4096)
// 	_, err := r.Read(buf)
// 	if err != nil {
// 		return err
// 	}
// 	lenKey := binary.LittleEndian.Uint32(buf[0:4])
// 	lenVal := binary.LittleEndian.Uint32(buf[4:8])

// 	ent.key = buf[8:8 + lenKey]
// 	ent.val = buf[8 + lenKey:8 + lenKey + lenVal]
// 	// fmt.Println("here")
// 	// fmt.Println(buf[8:8+lenKey])
// 	// fmt.Println(buf[8 + lenKey:8+lenKey+lenVal])
// 	return nil
// }

//more optimized version
func (ent *Entry) Decode(r io.Reader) error {
	var header [8]byte
	if _, err := io.ReadFull(r, header[:]); err != nil {
		return err
	}
	lenKey := binary.LittleEndian.Uint32(header[0:4])
	lenVal := binary.LittleEndian.Uint32(header[4:8])

	data := make([]byte, lenKey + lenVal)
	if _, err := io.ReadFull(r, data); err != nil {
		return err
	}
	ent.key = data[:lenKey]
	ent.val = data[lenKey:]
	return nil
}

// QzBQWVJJOUhU https://trialofcode.org/
