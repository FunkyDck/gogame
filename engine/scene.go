package engine

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Scene struct {
    camera *Camera
    terrain *Terrain
}

type Camera struct {
    pos mgl32.Vec3
    target mgl32.Vec3
    fovDegrees float32
    near float32
    far float32
}

func NewCamera() *Camera {
    return &Camera{
        pos: mgl32.Vec3{3, 3, 3},
        fovDegrees: 45,
        near: 0.1,
        far: 100,
    }
}

