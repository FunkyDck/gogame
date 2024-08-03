package engine

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"

    "dcbrwn.io/gogame/data"
)

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source + "\x00")
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func createProgram(
    vert string,
    frag string,
) (uint32, error) {
    vertexSource, err := data.ReadFile(vert)
    if err != nil {
        return 0, fmt.Errorf("failed to create shader: %w", err)
    }

	vertexShader, err := compileShader(string(vertexSource), gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

    fragmentSource, err := data.ReadFile(frag)
    if err != nil {
        return 0, fmt.Errorf("failed to create shader: %w", err)
    }

	fragmentShader, err := compileShader(string(fragmentSource), gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog, nil
}

