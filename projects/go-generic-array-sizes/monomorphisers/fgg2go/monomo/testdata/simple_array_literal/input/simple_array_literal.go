package main

type Arr[N const] [N]int

func main() {
	_ = Arr[2]{1, 2}
}
