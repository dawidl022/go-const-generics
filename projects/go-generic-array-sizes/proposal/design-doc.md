# Proposal: Const generics

Author(s): Dawid Lachowicz

Last updated: January 2024

Discussion at https://go.dev/issue/NNNNN.

## Abstract

This proposal aims to address one of the omissions stated in the [type parameter
proposal from
2021](https://go.googlesource.com/proposal/+/refs/heads/master/design/go2draft-type-parameters.md).

> No parameterization on non-type values such as constants. This arises most
> obviously for arrays, where it might sometimes be convenient to write type
> Matrix[n int] [n][n]float64. It might also sometimes be useful to specify
> significant values for a container type, such as a default value for elements.

There is an [existing proposal](https://github.com/golang/go/issues/44253) that
also aims to resolve this, however, it has been open for several years, and
this new proposal outlines a slightly different approach to the one in the
previous proposal.

This proposal would introduce a new predeclared (non-basic) type set interface
to the language, `const`, to be used as a type parameter constraint, intended to
be instantiated with constant non-negative integers, and one that can be used
to define the length of an array type.

## Background

In Go 1.18 type parameters were introduced to the language that allowed
programmers to genericise over types. All the built-in compound data types can
be fully parameterized with type parameters (slices: `[]T`, maps: `map[K]V` and
channels: `chan T`), except for arrays, so this feature would bridge that gap
with `[N]T`.

From the feedback on the [previous proposal aiming to tackle this
problem](https://github.com/golang/go/issues/44253), it is clear that many
people would benefit from introducing this feature into the language.
Specifically the feature would be aimed at more experienced users who have
justified use cases for using arrays over slices, and would enjoy the benefits
of reduced code duplication, or would like to expose some generic functionality
using underlying arrays as part of a library.

## Proposal

Conceptually `const` can be expressed as:

```go
type const interface {
    0 | 1 | 2 | 3 | ...
}
```

where `...` that the union contains all non-negative integers. Just other union
type set interfaces can be instantiated with one of the members of the union, a
type parameter bound by `const` can be instantiated with a constant expression
that evaluates to a non-negative integer at compile time.

The definition of const above is conceptual and would be implemented on the
language level, similar to the `comparable` interface.

Since `const` is a [non-basic
interface](https://go.dev/ref/spec#Interface_types), it can only be used to
constrain type parameters.

[A precise statement of the proposed change.]

## Rationale

[A discussion of alternate approaches and the trade offs, advantages, and disadvantages of the specified approach.]

## Compatibility

[A discussion of the change with regard to the
[compatibility guidelines](https://go.dev/doc/go1compat).]

## Implementation

[A description of the steps in the implementation, who will do them, and when.
This should include a discussion of how the work fits into [Go's release cycle](https://go.dev/wiki/Go-Release-Cycle).]

## Open issues (if applicable)

[A discussion of issues relating to this proposal for which the author does not
know the solution. This section may be omitted if there are none.]
