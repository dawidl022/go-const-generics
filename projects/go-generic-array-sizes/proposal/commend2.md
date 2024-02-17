## Syntax for general `const` types

It's clear from the comments that it's worth considering generalising the
proposal to other `const` types, not just for array lengths.

In which case, apart from the `const` keyword, the type of values should be
specifiable in the constraint, e.g. `uint8`, or `int64`.

Syntactically, we have a number of options, e.g.:

1. `chan T` style postfix: `const T`
    ```go
    func foo[N const int]() {
        // some code
    }
    ```

2.  `map[K]V` style postfix: `const[T]`
    ```go
    func foo[N const[int]] {
        // some code
    }
    ```

3. Rust style `const N T`:
   ```go
   func foo[const N int] {
        // some code
   }
   ```

My preference is for 1 or 2, because it plays better with using a single
constraint for multiple type parameters:

```go
type Matrix[N, M const[int]] [N][M]int
```

With option 3 this would be a bit awkward:

```go
type Matrix[const N, M int]
```

Another benefit of 1 and 2 over 3, is that `const T` or `const[T]` can be seen
as a type, just as `chan T` and `map[K]V` are types, as opposed to `const` being
some sort of type constraint modifier. This would play nicely if we extend types
to allow for value sets (i.e. unions of values), since they could simply be
named types and need no additional modifier.
