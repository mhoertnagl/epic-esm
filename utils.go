package main

import "bufio"

// func WriteArrayInt32BigEndian(w *bufio.Writer, i []uint32) {
// 	for _, code := range i {
// 		WriteInt32BigEndian(w, code)
// 	}
// }

func WriteInt32BigEndian(w *bufio.Writer, i uint32) {
	w.WriteByte(byte(i >> 24))
	w.WriteByte(byte(i >> 16))
	w.WriteByte(byte(i >> 8))
	w.WriteByte(byte(i))
}
