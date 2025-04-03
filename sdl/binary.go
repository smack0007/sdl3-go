package sdl

import (
	"encoding/binary"
	"math"
	"unsafe"
)

func readBool[T any](data *T, offset int) bool {
	return readUint8(data, offset) != 0
}

func readBytes[T any](data *T, offset int, length int) []byte {
	return unsafe.Slice((*byte)(unsafe.Add((unsafe.Pointer)(data), offset)), length)
}

func readFloat32[T any](data *T, offset int) float32 {
	return math.Float32frombits(readUint32(data, offset))
}

func readInt32[T any](data *T, offset int) int32 {
	// TODO: Not sure this is correct.
	return (int32)(readUint32(data, offset))
}

func readUint8[T any](data *T, offset int) uint8 {
	return *(*uint8)(unsafe.Add((unsafe.Pointer)(data), offset))
}

func readUint16[T any](data *T, offset int) uint16 {
	return binary.NativeEndian.Uint16(readBytes(data, offset, 2))
}

func readUint32[T any](data *T, offset int) uint32 {
	return binary.NativeEndian.Uint32(readBytes(data, offset, 4))
}
