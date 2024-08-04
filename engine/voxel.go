package engine

import (
    "math"
)

const Void uint16 = 0

type VoxelChunk struct {
    valid bool
    voxels []uint16
    vertices []float32
    size i32vec3
    numBlocks uint32
}

func NewVoxelChunk(size i32vec3) *VoxelChunk {
    numBlocks := size.X * size.Y * size.Z
    return &VoxelChunk {
        voxels: make([]uint16, numBlocks),
        vertices: make([]float32, numBlocks * 4), 
        size: size,
    }
}

func (vox *VoxelChunk) toIndex(p i32vec3) int32 {
    layerSize := vox.size.X * vox.size.Z
    return p.X + p.Z * vox.size.Z + p.Y * layerSize
}

func (vox *VoxelChunk) toPos(index int32) i32vec3 {
    zSize := float64(vox.size.Z)
    layerSize := float64(vox.size.X) * zSize
    fIndex := float64(index)

    y := math.Floor(fIndex / layerSize)
    l := fIndex - (y * layerSize)
    
    return i32vec3{
        X: int32(math.Mod(l, zSize)),
        Z: int32(math.Floor(l / zSize)),
        Y: int32(y),
    }
}

func (vox *VoxelChunk) Put(p i32vec3, value uint16) {
    index := vox.toIndex(p)
    old := vox.voxels[index]
    vox.voxels[index] = value
    if old == Void && value != Void {
        vox.numBlocks += 1
    } else if old != Void && value == Void {
        vox.numBlocks -= 1
    }
    vox.valid = false
}

func (vox *VoxelChunk) Get(p i32vec3) uint16 {
    return vox.voxels[vox.toIndex(p)]
}

func (vox *VoxelChunk) GetVertices() []float32 {
    if !vox.valid {
        vox.regen()
    }

    return vox.vertices
}

var cube = []float32{
    0.0, 1.0, 1.0,
    1.0, 1.0, 1.0,
    0.0, 0.0, 1.0,
    1.0, 0.0, 1.0,
    1.0, 0.0, 0.0,
    1.0, 1.0, 1.0,
    1.0, 1.0, 0.0,
    0.0, 1.0, 1.0,
    0.0, 1.0, 0.0,
    0.0, 0.0, 1.0,
    0.0, 0.0, 0.0,
    1.0, 0.0, 0.0,
    0.0, 1.0, 0.0,
    1.0, 1.0, 0.0,
}

func (vox *VoxelChunk) regen() error {
    vox.vertices = vox.vertices[:(int(vox.numBlocks) * 3)]
    cursor := 0

    for y := int32(0); y < vox.size.Y; y += 1 {
        for z := int32(0); z < vox.size.Z; z += 1 {
            for x := int32(0); x < vox.size.X; x += 1 {
                b := vox.Get(i32vec3{x, y, z})
                if b == Void {
                    continue
                }

                vox.vertices[cursor + 0] = float32(x)
                vox.vertices[cursor + 1] = float32(y)
                vox.vertices[cursor + 2] = float32(z)
                    cursor += 3
            }
        }
    }

    vox.valid = true
    return nil
}

