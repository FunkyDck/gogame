package engine

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Renderer struct{
    uniforms struct {
        Projection mgl32.Mat4
        Camera mgl32.Mat4
        Model mgl32.Mat4
    }

    terrainVAO uint32
    terrainProgram uint32
}

func NewRenderer() (*Renderer, error) {
    r := &Renderer {}

    program, err := createProgram(
        "shaders/terrain/terrain.vert",
        "shaders/terrain/terrain.frag",
        "shaders/terrain/terrain.geom",
    )
    if err != nil {
        return nil, err
    }
    r.terrainProgram = program

    return r, nil
}

func (r *Renderer) ApplyUniforms(program uint32) {
    u := r.uniforms

	engTimeUniform := gl.GetUniformLocation(program, gl.Str("engineTime\x00"))
	gl.Uniform1d(engTimeUniform, glfw.GetTime())

	projectionUniform := gl.GetUniformLocation(program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &u.Projection[0])

	cameraUniform := gl.GetUniformLocation(program, gl.Str("camera\x00"))
	gl.UniformMatrix4fv(cameraUniform, 1, false, &u.Camera[0])

	modelUniform := gl.GetUniformLocation(program, gl.Str("model\x00"))
	gl.UniformMatrix4fv(modelUniform, 1, false, &u.Model[0])

	// textureUniform := gl.GetUniformLocation(program, gl.Str("tex\x00"))
	// gl.Uniform1i(textureUniform, 0)

	gl.BindFragDataLocation(program, 0, gl.Str("outputColor\x00"))
}

func (r *Renderer) Render(
    scene *Scene,
    config *Config,
) {
    gl.ClearColor(0.0, 0.0, 0.0, 1.0)
    gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
    gl.Enable(gl.DEPTH_TEST)
    gl.Disable(gl.CULL_FACE)
    gl.DepthFunc(gl.LESS)

    r.uniforms.Projection = mgl32.Perspective(
        mgl32.DegToRad(scene.camera.fovDegrees),
        float32(config.Window.Width) / float32(config.Window.Height),
        scene.camera.near,
        scene.camera.far)
    
    camPos := mgl32.Vec3{3, 3, 3}
    camPos = mgl32.Rotate3DY(float32(glfw.GetTime())).Mul3x1(camPos)

    r.uniforms.Camera = mgl32.LookAtV(
        camPos,
        scene.camera.target,
        mgl32.Vec3{0, 1, 0})

    r.uniforms.Model = mgl32.Ident4()

    wasValid := scene.terrain.IsValid()
    vertices := scene.terrain.GetVertices()

    if !wasValid {
        r.terrainVAO = makeVao(vertices)
    }

    gl.UseProgram(r.terrainProgram)
    r.ApplyUniforms(r.terrainProgram)

	gl.BindVertexArray(r.terrainVAO)
    gl.PointSize(5.0)
	gl.DrawArrays(gl.POINTS, 0, int32(len(vertices)))
}

func makeVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 42*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

