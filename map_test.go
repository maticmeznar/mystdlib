package mystdlib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapKeyExists(t *testing.T) {
	m := map[string]interface{}{
		"key1": "value1",
	}
	assert.True(t, MapKeyExists(m, "key1"))
	assert.False(t, MapKeyExists(m, "key2"))
}
