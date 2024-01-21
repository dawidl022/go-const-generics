package reversed

type arrayGeneric[T any] interface {
	[0]T |
	[1]T |
	[2]T |
	[4]T |
	[8]T |
	[16]T |
	[32]T |
	[64]T |
	[128]T |
	[256]T |
	[512]T |
	[1024]T |
	[2048]T |
	[4096]T |
	[8192]T |
	[16384]T |
	[32768]T |
	[65536]T |
	[131072]T |
	[262144]T |
	[524288]T |
	[1048576]T |
	[2097152]T |
	[4194304]T |
	[8388608]T |
	[16777216]T |
	[33554432]T
}

func reversedArrayGenericUnion[T any, A arrayGeneric[T]](arr A) A {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
	return arr
}
