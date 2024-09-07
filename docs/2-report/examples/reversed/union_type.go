package reversed

type array[T any] interface {
	[2]T | [3]T | [4]T | [5]T
}

func reversedArray[T any, A array[T]](arr A) A {
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
	return arr
}
