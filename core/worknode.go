package core

type WorkNode[P comparable] struct {
	work Pair[P, Work]
}
