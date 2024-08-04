package engine

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestGeneratedVertices(t *testing.T) {
    v := NewVoxelChunk(i32vec3{3, 3, 3})
    v.Put(i32vec3{0, 1, 2}, 1)
    v.Put(i32vec3{1, 0, 0}, 1)
    actual := v.GetVertices()
    expected := []float32{
        1.0, 0.0, 0.0,
        0.0, 1.0, 2.0,
    }

    assert.Equal(t, expected, actual, "")
}
