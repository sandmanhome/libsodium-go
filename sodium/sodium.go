package sodium

import "fmt"
import "unsafe"

// #include <stdio.h>
// #include <sodium.h>
import "C"

func Init() {
	result := int(C.sodium_init())
	if result != 0 {
		panic(fmt.Sprintf("Sodium initialization failed, result code %d.",
			result))
	}
}

func RuntimeGetCpuFeatures() int {
	return int(C.sodium_runtime_get_cpu_features())
}

func RuntimeHasNeon() bool {
	return C.sodium_runtime_has_neon() != 0
}

func RuntimeHasSse2() bool {
	return C.sodium_runtime_has_sse2() != 0
}

func RuntimeHasSse3() bool {
	return C.sodium_runtime_has_sse3() != 0
}

func MemZero(buff1 []byte) {
	if len(buff1) > 0 {
		C.sodium_memzero(unsafe.Pointer(&buff1[0]), C.size_t(len(buff1)))
	}
}

func MemCmp(buff1, buff2 []byte, length int) int {
	if length >= len(buff1) || length >= len(buff2) {
		panic(fmt.Sprintf("Attempt to compare more bytes (%d) than provided " + 
			"(%d, %d)", length, len(buff1), len(buff2)))
	}
	return int(C.sodium_memcmp(unsafe.Pointer(&buff1[0]),
		unsafe.Pointer(&buff2[0]),
		C.size_t(length)))
}

func Bin2hex(bin []byte) string {
	maxlen := len(bin) * 2
	binPtr := (*C.uchar)(unsafe.Pointer(&bin[0]))
	buf := (*C.char)(C.malloc(C.size_t(maxlen)))
	defer C.free(unsafe.Pointer(buf))
	C.sodium_bin2hex(buf, C.size_t(maxlen), binPtr, C.size_t(len(bin)))

	return C.GoString(buf)
}
