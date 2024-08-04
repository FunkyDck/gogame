package engine

type Terrain struct {
    chunk *VoxelChunk
}

func NewTerrain() (*Terrain, error) {
    terrain := &Terrain{
        chunk: NewVoxelChunk(i32vec3{16, 256, 16}),
    }

    return terrain, nil
}

func (t *Terrain) Put(p i32vec3, value uint16) {
    t.chunk.Put(p, value)
}

func (t *Terrain) IsValid() bool {
    return t.chunk.valid
}

func (t *Terrain) GetVertices() []float32 {
    return t.chunk.GetVertices()
}
