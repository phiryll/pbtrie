package btrie

import "iter"

type adjFunction[T any] func([]T) iter.Seq[T]

// The returned sequence references a volatile internal slice,
// clone it if you need it after a step in the iteration.
func preOrder[T any](root T, adj adjFunction[T]) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		preOrderRecurse([]T{root}, adj, yield)
	}
}

// Returns a value of true if done (some yield has returned false).
func preOrderRecurse[T any](path []T, adj adjFunction[T], yield func([]T) bool) bool {
	if !yield(path) {
		return true
	}
	for node := range adj(path) {
		if preOrderRecurse(append(path, node), adj, yield) {
			return true
		}
	}
	return false
}
