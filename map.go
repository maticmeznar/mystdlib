package mystdlib

// MapKeyExists returns true if the map key exists. False otherwise. Implemented using generics.
func MapKeyExists[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}
