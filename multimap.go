// Package multimap provides an abstract MultiMap interface.
//
// Multimap is a collection that maps keys to values, similar to map.
// However, each key may be associated with multiple values.
//
// You can visualize the contents of a multimap either as a map from keys to nonempty collections of values:
//    - a --> 1, 2
//    - b --> 3
// ... or a single "flattened" collection of key-value pairs.
//    - a --> 1
//    - a --> 2
//    - b --> 3
//
// Similar to a map, operations associated with this data type allow:
// - the addition of a pair to the collection
// - the removal of a pair from the collection
// - the lookup of a value associated with a particular key
// - the lookup whether a key, value or key/value pair exists in this data type.
package multimap

// Entry represents a key/value pair inside a multimap.
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

// MultiMap interface that all multimaps implement.
type MultiMap[K comparable, V any] interface {
	Get(key K) (values []V, found bool)

	Put(key K, value V)
	PutAll(key K, values []V)

	Remove(key K, value V)
	RemoveAll(key K)

	Contains(key K, value V) bool
	ContainsKey(key K) bool
	ContainsValue(value V) bool

	Entries() []Entry[K, V]
	Keys() []K
	KeySet() []K
	Values() []V

	Clear()
	Empty() bool
	Size() int
}
