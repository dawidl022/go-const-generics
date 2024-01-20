package reversed

type array interface {
	[0]int |
	[1]int |
	[2]int |
	[4]int |
	[8]int |
	[16]int |
	[32]int |
	[64]int |
	[128]int |
	[256]int |
	[512]int |
	[1024]int |
	[2048]int |
	[4096]int |
	[8192]int |
	[16384]int |
	[32768]int |
	[65536]int |
	[131072]int |
	[262144]int |
	[524288]int |
	[1048576]int |
	[2097152]int |
	[4194304]int |
	[8388608]int |
	[16777216]int |
	[33554432]int
}

func reversedArrayUnion[A array](arr A) A {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
	return arr
}
