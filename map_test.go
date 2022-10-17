package mystdlib

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMapKeyExists(t *testing.T) {
	m1 := map[string]int{
		"key1": 42,
	}
	assert.True(t, MapKeyExists(m1, "key1"))
	assert.False(t, MapKeyExists(m1, "key2"))

	id := uuid.New()
	m2 := map[uuid.UUID]string{
		id: "good",
	}
	assert.True(t, MapKeyExists(m2, id))
	assert.False(t, MapKeyExists(m2, uuid.New()))
}
