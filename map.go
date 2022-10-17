package mystdlib

// MapKeyExists returns true if the map key exists. False otherwise.
// Currently broken
func MapKeyExists(m interface{}, key interface{}) bool {
	m1 := m.(map[interface{}]interface{})
	_, ok := m1[key]
	if !ok {
		return false
	} else {
		return true
	}
}
