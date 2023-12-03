# Generically sized arrays in Go

Student number: 210150908

Dawid Lachowicz

Supervisor: Raymond Hu

<!-- TODO insert generic gopher -->

---

## Background: Arrays

How various languages treat arrays

<!-- TODO insert gophers inside an array -->

---

## Problem: Size of an array in Go is part of its type

<!-- TODO insert oh-no gopher  -->

<!-- TODO insert code snippet with hard-coded array sizes -->

```go
func reversed(arr [5]int) [5]int {
	n := len(arr)
	for i := 0; i < n / 2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
	return arr
}
```

---

## Literature Review: Featherweight Go

Introducing generics to Go (Griesemer et al., 2020)

```go
func reversed[T any](arr [5]T) [5]T {
	n := len(arr)
	for i := 0; i < n / 2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
	return arr
}
```

---

## Solution: Numerical type parameters

Type parameters for array sizes

```go
func reversed[T any, N const](arr [N]T) [N]T {
	n := len(arr)
	for i := 0; i < n / 2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
	return arr
}
```

---

## Rust: const generics

Rust already tackled this issue

```rust
struct ArrayPair<T, const N: usize> {
    left: [T; N],
    right: [T; N],
}
```

(The const generics project group, 2021)

<!-- TODO insert rust snippet -->

<!-- TODO insert rustling -->

---

## Methodology: Formal rules

Syntax, reduction and typing rules for Featherweight Go with arrays

<!-- TODO include snippet of a single rule for each of the 3 types -->


<div class="flex-container">


<div class="flex-item">

Syntax

$T ::= [n]t$ 

</div>
<div class="flex-item">

Reduction

`$$
\begin{array}{c}
   n \in indexBounds(t_A) \\ \hline
   t_A\{\overline{v}\}[n] \to v_n \\
\end{array}
$$`

</div>
<div class="flex-item">

Typing

`$$
\begin{array}{c}
    n \ge 0~~~~~t~~ok\\ \hline
    [n]t~~ok
\end{array}
$$`

</div>
---

## Methodology: Interpreter implementation

```
reduction step 1: Arr{4, 6}[Foo{3, Arr{1, 2}}.y.first()]
program well typed

reduction step 2: Arr{4, 6}[Arr{1, 2}.first()]
program well typed

reduction step 3: Arr{4, 6}[Arr{1, 2}[0]]
program well typed

reduction step 4: Arr{4, 6}[1]
program well typed

reduction step 5: 6
program well typed
```

---

## Next steps: Monomorphiser implementation

<div class="flex-item">

```go
func reversed[T any, N const](arr [N]T) [N]T {
    // reversing code...
}
func main() {
    fmt.Println(reversed([5]int{1, 2, 3, 4, 5}))
}
```
</div>

<div class="flex-item">
$\downarrow$
</div>

<div class="flex-item">

```go
func reversed_5[T any](arr [5]T) [5]T {
    // reversing code...
}
func main() {
    fmt.Println(reversed_5([5]int{1, 2, 3, 4, 5}))
}
```

</div>

</div>

---

## Next steps: Proposal + compiler implementation

<img
src="https://raw.githubusercontent.com/primer/octicons/main/icons/git-pull-request-16.svg"
height="200"
style="filter: invert(0.9);">

---

## Questions

References:

<small>

Griesemer, R., Hu, R., Kokke, W., Lange, J., Taylor, I.L., Toninho, B., Wadler,
P., and Yoshida, N., 2020. Featherweight Go. Proc. ACM Program. Lang. [Online],
4(OOPSLA). Available from: https://doi.org/10.1145/3428217.

The const generics project group, 2021. Const generics MVP hits beta! [Online].
Rust Blog. Available from:
https://blog.rust-lang.org/2021/02/26/const-generics-mvp-beta.html [Accessed
November 12, 2023].

</small>


Images:

<small>

https://github.com/primer/octicons

</small>
