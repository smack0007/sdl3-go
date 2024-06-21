package sdl

import (
	"encoding/binary"
)

func readBytes(data []byte, offset int, length int) []byte {
	return data[offset : offset+length]
}

func readInt32(data []byte, offset int) int32 {
	// TODO: Not sure this is correct.
	return int32(binary.NativeEndian.Uint32(data[offset : offset+4]))
}

func readUint8(data []byte, offset int) uint8 {
	return data[offset]
}

func readUint16(data []byte, offset int) uint16 {
	return binary.NativeEndian.Uint16(data[offset : offset+2])
}

func readUint32(data []byte, offset int) uint32 {
	return binary.NativeEndian.Uint32(data[offset : offset+4])
}
