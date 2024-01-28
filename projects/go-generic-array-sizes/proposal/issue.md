https://github.com/golang/go/issues/44253 proposes to introduce generic
parameterization of array sizes. This issue has been open for 3 years, with
seemingly no recent progress.

The new proposal differs in that it introduces an explicit numerical "const"
type parameter, and restricts the kinds of expressions that can be used to
instantiate a const type parameter to existing constant expressions, and lone
"const" type parameters, in the [same fashion as
Rust](https://blog.rust-lang.org/2021/02/26/const-generics-mvp-beta.html#no-complex-generic-expressions-in-const-arguments).

Having this restriction in place limits the problems const generics can solve,
but it also eliminates a lot of complexity associated with implementing const
generics, that was discussed in the issue of the previous proposal.

If the community is happy to proceed with this proposal, I would like to have a
go at implementing the feature. As such, any concerns regarding potential
implementation difficulties would be appreciated!

Detailed design document: (TODO)

## Proposal

The basis of this proposal is to introduce a new set of types that can be used
as type arguments. This new type set would be made up of constant non-negative
integers. Since interfaces represent type sets in Go, we can (conceptually)
define such as type as:

```go
type const interface {
    0 | 1 | 2 | 3 | ...
}
```

where the `...` denotes that the pattern continues for all non-negative integers
supported by Go. The `const` interface can be used to constrain a type
parameter, and since it is a [non-basic
interface](https://go.dev/ref/spec#Interface_types) it cannot be used as the
type of a variable (we can already use `const` declarations to hold values, but
they are broader than this proposed interface).

Array types accept two "type" parameters - the length and the element type. The
bound of the length parameter is exactly the `const` interface. This proposal
generalises the `const` interface to be applicable to user-defined type
parameters - those we know from generics.

Since all type parameters also implement their own constraints (i.e. a type
parameter constrained by `T` can be used as a type argument constrained by the
same type, or a supertype of `T`), it means that `const` type parameters can be
used to instantiate other `const` type parameters, including array lengths.

Expressions involving type parameters, other than a lone `const`
type parameter `N`, cannot be used as `const` type arguments. This approach is
used by
[Rust](https://blog.rust-lang.org/2021/02/26/const-generics-mvp-beta.html), and
can make the implementation much more feasible (discussed in detail in the
design document). E.g. the following would not be allowed:

```go
func foo[N const]() {
    // compile-time error: the expression `N + 1` cannot be used as a type argument
    foo[N + 1]()
}
```

In addition, despite `const` type parameters implementing `const` themselves,
when used as a value, as opposed to a type parameter, they would evaluate to a
non-constant `int`. E.g. the following would not be allowed:

```go
func foo[N const]() {
    // compile-time error: type parameter N cannot be used as constant value
    const n = N
}
```

This is to avoid `n` being used as part of another constant expression that is
then used as a `const` type argument, in effect enforcing the previous
limitation of no "no complex generic expressions in const arguments".

The above restriction also resolves the issue that was heavily discussed in
https://github.com/golang/go/issues/44253#issuecomment-821047754 of exploiting
`const` type parameters to perform complex compile time computation, such as SAT
solving.

<!-- TODO Please describe as precisely as possible the change to the language. -->

<!-- TODO Please also describe the change informally, as in a class teaching Go. -->

## Examples

One of the simplest useful examples is declaring a square matrix type:

```go
type Matrix[T any, N const] [N][N]T
```

which can then be instantiated as follows:

```go
myMatrix := Matrix[int, 2]{{1, 2}, {3, 4}}
```

Of course, `const` type parameters can also be used directly in functions:

```go
func reversed[T any, N const](arr [N]T) [N]T {
    for i := 0; i < N/2; i++ {
        arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
    }
    return arr
}
```

<!-- TODO Show example code before and after the change. -->

## Changes to language spec

[Array types](https://go.dev/ref/spec#Array_types) would also accept
`TypeName`s, which (in my understanding) type parameters fall under.

```
ArrayLength = Expression | TypeName .
```

[Types](https://go.dev/ref/spec#Types) would also include constant expressions,
more precisely constant expressions that evaluate to non-negative integers. They
would be used to instantiate `const` type parameters.

```
Type = TypeName [ TypeArgs ] | TypeLit | Expression | "(" Type ")" .
```

[Type constraints](https://go.dev/ref/spec#TypeConstraint) would accept the
keyword `const` to denote the interface containing all non-negative integer
types.

```
TypeConstraint = TypeElem | "const" .
```

<!-- ## Comparison with previous proposal

This proposal introduces a cleaner, and more intuitive way of expressing that a
type is parameterised by an integer. The `const` type parameter is explicit,
which is consistent with the existing generics system in Go. -->

## Other questions from template

Q: Who does this proposal help, and why?

This proposal is aimed at experienced Go users, who have justified use cases for
using arrays over slices, and would enjoy the benefits of reduced code
duplication with const generics.

Q: Is this change backward compatible?

Yes, this change is backwards compatible. It is not possible to name a regular
interface the keyword `const`, so no existing code would break.

Q: What is the cost of this proposal? (Every language change has a cost).

- What is the compile time cost?

The cost should be approximately the same as the existing cost of compiling
generic code.

- What is the run time cost?

This depends on how the feature will be implemented. Using a full
monomorphisation model, there would be no runtime costs over using non-generic
arrays. Using the [GC shape
stencilling](https://github.com/golang/proposal/blob/master/design/generics-implementation-dictionaries-go1.18.md)
approach would likely have the same runtime cost as existing generic code.

Q: Can you describe a possible implementation? Do you have a prototype? (This is
not required.)

I am working on a paper in the style of [Featherweight
Go](https://dl.acm.org/doi/10.1145/3428217), that formalises arrays and
numerical type parameters, that would be implemented as part of this proposal.
It will come with prototype interpreters for the new language feature in the
style of https://github.com/rhu1/fgg (i.e. for a subset of Go). I hope to make
this publicly available, sometime post May this year.

Q: Would this change make Go easier or harder to learn, and why?

Seeing that this is an additional language feature, it would make Go more
difficult to learn simply because there is one feature more. The rest of the
language will be unaffected in terms of easiness to learn.

Q: Orthogonality: how does this change interact or overlap with existing
features?

Explicit numerical type parameters would make generic arrays a first-class
feature of Go, consistent with the rest of the language. All the existing
compound data types in Go can already be fully type parameterised (slices:
`[]T`, maps: `map[K]V` and channels: `chan T`), except for arrays, so this
feature would bridge that gap with `[N]T`.

Numerical type parameters would also become first class. Currently, arrays are
the only type which accept a numerical type parameter, to parameterize the
length of an array type. The `const` interface would allow any type or function
to accept a constant integer (or another `const` bounded type parameter) as a
type argument.

Q: Would you consider yourself a novice, intermediate, or experienced Go
programmer?

Experienced.

Q: What other languages do you have experience with?

Rust, Kotlin, Java, TypeScript, Python
