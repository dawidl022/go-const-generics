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

3. Rust style: `const N T`
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

`const T` or `const[T]` can be seen as a type, just as `chan T` and `map[K]V`
are types, as opposed to `const` being some sort of type constraint modifier in
option 3.

## Array length type

According to the spec, the length of an array type:

> must evaluate to a non-negative constant representable by a value of type
> `int` 

This means that neither `int` not `uint` captures the valid values for array
lengths, and so in theory neither `const[int]` nor `const[uint]` could be used
to constrain a numerical type parameter used to size an array, since the
programmer could instantiate such a type parameter with:

1. a negative value for `const[int]`
2. a value larger than `math.MaxInt` for `const[uint]`

We could either loosen the restrictions and allow arrays to be sized by `uint`
(not sure how this would affect things internally), or introduce a new way of
expressing such a type, e.g. `aint` (array int).

Then `const[aint]` could be used to denote the type parameter can be used to
size an array. Alternatively (or in addition to), we could have a lone `const`
be the shorthand constraint for `const[aint]`.

## Types with underlying `const` types

To address the issue of using types such as `time.Duration` as a `const` type
parameter, we have a couple of options (or we can use all of them):

1. Allow types which have underlying types that can be `const` to be used as
   `const` constraints, e.g. `const[time.Duration]`. Such a constraint can only
   be instantiated with constants of type `time.Duration` (or untyped
   constants).
2. Allow type set `const` constraints, e.g. `const[~int64]`. Such a constraint
   can be instantiated with any typed constant whose underlying type is `int64`
   or an untyped constant.
3. Use conversions where necessary to convert `const` type parameters to other
   types, e.g. `time.Duration(N)` where `N` is a type parameter bound by
   `const[int64]`. If the underlying type of the target type matches with the
   `const` constraint type (`int64` in the example), then this conversion can be
   done at compile time. This avoids the problem of determining whether a type
   conversion can be allowed at compile time, discussed in
   https://github.com/golang/go/issues/44253#issuecomment-821047754.
