package main

type Arr[N const] [N]int

func main() {
	_ = Arr[2]{1, 2}[0] + Arr[3]{1, 2, 3}[1]
}
