package engine

import (
	"fmt"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"

    "dcbrwn.io/gogame/scripting"
)

type Engine struct {
    config *Config
    scene *Scene
    window *glfw.Window
}

func NewEngine(configpath string) (*Engine, error) {
    config, err := LoadConfig(configpath)
    if err != nil {
        return nil, fmt.Errorf("failed to read config: %w", err)
    }

    engine := &Engine{
        config: config,
    }

    return engine, nil
}

func (e *Engine) Run() error {
    log.Printf("Starting engine...")
	runtime.LockOSThread()

    window, err := e.initGlfw()
    if err != nil {
        return err
    }

	defer glfw.Terminate()

	if err := gl.Init(); err != nil {
		return err
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

    go func() {
        _, err := scripting.Load(
            "scripts/main.py",
            "main script",
            nil,
        )
        if err != nil {
            panic(err)
        }
    }()

    e.scene = &Scene{}
    terrain, err := NewTerrain()
    e.scene.root = &Composite{
        Children: []Drawable{
            terrain,
        },
    }

	for !window.ShouldClose() {
        gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
        e.scene.Draw()
        glfw.PollEvents()
        window.SwapBuffers()
	}

    return nil
}

func (e *Engine) initGlfw() (*glfw.Window, error) {
	if err := glfw.Init(); err != nil {
        return nil, err
	}
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(
        e.config.Window.Width,
        e.config.Window.Height,
        "Go game!",
        nil,
        nil,
    )
	if err != nil {
        return nil, err
	}
	window.MakeContextCurrent()

	return window, nil
}
