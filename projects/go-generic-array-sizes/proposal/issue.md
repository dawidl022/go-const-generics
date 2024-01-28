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

Please describe as precisely as possible the change to the language.


<!-- TODO do we allow interfaces that are a union of expressions? -->


Please also describe the change informally, as in a class teaching Go.

## Examples

Show example code before and after the change.

<!-- TODO -->

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
