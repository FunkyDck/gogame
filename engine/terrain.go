package engine

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Terrain struct {
    glProgram uint32
    glVertexArray uint32 
    
    numVertices int32
}

var _ Drawable = &Terrain{}

func NewTerrain() (*Terrain, error) {
    terrain := &Terrain{}

    prog, err := createProgram(
        "shaders/basic.vert",
        "shaders/basic.frag",
    )
    if err != nil {
        return nil, err
    }

    terrain.glProgram = prog
    terrain.glVertexArray = makeVao(triangle)
    terrain.numVertices = 3

    return terrain, nil
}

func (t *Terrain) Draw() {
	gl.UseProgram(t.glProgram)

	gl.BindVertexArray(t.glVertexArray)
	gl.DrawArrays(gl.TRIANGLES, 0, t.numVertices)
}

var (
	triangle = []float32{
		0, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
	}
)

func makeVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}
