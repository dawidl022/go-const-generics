type werner

type Dim interface {
    [...]struct{}
}

type Matrix[D Dim] [len(D)][len(D)]int
