package reversed

func reversed[T any, N const](arr [N]T) [N]T {
    n := len(arr)
    for i := 0; i < n/2; i++ {
        arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
    }
    return arr
}
