package mystdlib

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestIsInSlice(t *testing.T) {
	assert.True(t, IsInSlice("hello", []string{"hello", "burek"}))
	assert.False(t, IsInSlice("lala", []string{"hello", "burek"}))

	assert.True(t, IsInSlice(42, []int{1, 4, 42, 54}))
	assert.False(t, IsInSlice(99, []int{1, 4, 42, 54}))

	id := uuid.MustParse("d5e45b39-9731-4b76-8929-09c3a4c3759d")
	assert.True(t, IsInSlice(id, []uuid.UUID{uuid.New(), uuid.New(), id, uuid.New()}))
	assert.False(t, IsInSlice(id, []uuid.UUID{uuid.New(), uuid.New(), uuid.New()}))
}

func BenchmarkIsInSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsInSlice("match", []string{"hello", "welcome", "dreams", "joy", "keyboard", "table", "match", "door"})
	}
}
