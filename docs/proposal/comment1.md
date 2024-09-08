> I'm concerned that using `const` as the constraint will make it difficult for
> us to add any other type of numeric constraint. It's not clear that we would
> ever want any such constraint. But if we ever do, it's not going to be able to
> look like `const`.

One potential way of expressing other numerical constraints could be to make
`const` itself generic. E.g. we could have `const[uint8]`, `const[float64]`,
that would accept numerical type arguments from the range of the type `const` is
instantiated with.

As far as I'm aware, there is no explicit numerical type for array lengths. The
[spec](https://go.dev/ref/spec#Array_types) says:

> it must evaluate to a **non-negative constant** representable by a value of
> type `int`

so having the lone `const` constraint represent that concept seems reasonable.

Another possibility for creating other types of numerical constraints, could be
to set some explicit bounds on the constraint, e.g. using the familiar slicing
notation. `const[1:]` would represent any numerical (`int`) type parameter
greater than equal to 1. This could be used to e.g. define the allowed
compile-time checked operations on an array of such a length.

In the current proposal, indexing a generic array by a constant is not allowed,
since the code should be valid for all instantiations. E.g.

```go
func first[N const](arr [N]int) {
    return arr[0] // compile-time error
}
```

would not compile, since `[0]int` is a valid instantiation of `[N]int`. However,
if we constrain the type parameter to `const[1:]`, then this operation becomes
valid, since the function `first` now only accepts arrays of at least length 1:

```go
func first[N const[1:]](arr [N]int) {
    return arr[0] // allowed and safe (no runtime panics can occur)
}
```

Explicit lower and upper bounds could also be used to loosen some of the
restrictions imposed by this proposal, namely of using constant expressions that
involve `const` type parameters and other constants as type arguments, to make
the examples presented in the https://github.com/golang/go/issues/44253 proposal
possible.

We can of course combine the two concepts, and define the constraint that can
be used for any array lengths as the spec defines it:

```go
func foo[N const[int][0:]](arr [N]int) {
    // some code
}
```

But since `const[int][0:]` seems like a common use case, using `const` as a
shorthand for the same type would be convenient anyways.

Of course whether or not we'd like to introduce such elaborate numerical type
constraints is a topic for a future discussion, but I believe introducing a
simple `const` constraint to begin with would not make defining the elaborate
constraints building on top of `const` impossible.

>  We should also bear in mind that there are other kinds of constraints that we
>  can't express, like saying that one type parameter must be assignable to
>  another.

Do I understand correctly, that by "one type parameter must be assignable to
another" we mean that one must *implement* the other? I believe this could be
expressed by making one type parameter the constraint of the other (which is not
currently allowed in Go).

```go
func foo[B any, A B](a A) {
    var myVar B = a
    // some more code
}
```

> We want to make sure that the "constraint language" is consistent and has good
> support for everything we might want to do in the future.

Is there a list I can check out of things we might want to express using type
parameter constraints in the future? I would be happy to have a look and
consider how this might affect the design of `const` generics. So far, having
type-set interfaces seems to work well, and we already have "special"
constraints for built-in operations such as `comparable`. So we can either view
`const` as in the same bucket as `comparable`, or as type-set interfaces being
extended with "value sets" (such as numerical values of specific types or
subranges of those types, or even individual values such as `0 | 1`, similar to
[literal types in
TypeScript](https://www.typescriptlang.org/docs/handbook/2/everyday-types.html#literal-types)).
Value sets could also be used as enum types.
