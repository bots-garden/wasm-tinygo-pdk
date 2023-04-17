package plugin

import (
	"unsafe"
)

const isFailure = rune('F')
const isSuccess = rune('S')

var handleFunction func(param []byte) ([]byte, error)

func SetHandle(function func(param []byte) ([]byte, error)) {
	handleFunction = function
}

//export callHandle
func callHandle(subjectPosition *uint32, length int) uint64 {
	subjectBytes := read(subjectPosition, length)
	retValue, err := handleFunction(subjectBytes)

	if err != nil {
		return failure([]byte(err.Error()))
	} else {
		// first byte == 82
		return success(retValue)
	}
}

// readBufferFromMemory returns a buffer from WebAssembly
func readBufferFromMemory(bufferPosition *uint32, length int) []byte {
	subjectBuffer := make([]byte, length)
	pointer := uintptr(unsafe.Pointer(bufferPosition))
	for i := 0; i < length; i++ {
		s := *(*int32)(unsafe.Pointer(pointer + uintptr(i)))
		subjectBuffer[i] = byte(s)
	}
	return subjectBuffer
}

func copyBufferToMemory(buffer []byte) uint64 {
	bufferPtr := &buffer[0]
	unsafePtr := uintptr(unsafe.Pointer(bufferPtr))

	ptr := uint32(unsafePtr)
	size := uint32(len(buffer))

	return (uint64(ptr) << uint64(32)) | uint64(size)
}

func read(bufferPosition *uint32, length int) []byte {
	return readBufferFromMemory(bufferPosition, length)
}

func success(buffer []byte) uint64 {
	return copyBufferToMemory(append([]byte(string(isSuccess)), buffer...))
}

func failure(buffer []byte) uint64 {
	return copyBufferToMemory(append([]byte(string(isFailure)), buffer...))
}

