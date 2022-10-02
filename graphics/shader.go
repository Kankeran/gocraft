package graphics

import (
	"embed"
	"fmt"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

//go:embed shaders
var shaderResources embed.FS

type ShaderProgram uint32

func NewShader(shaderName string) (*ShaderProgram, error) {
	vertexShader, err := compileShader(shaderName+".vs", gl.VERTEX_SHADER)
	if err != nil {
		return nil, err
	}

	fragmentShader, err := compileShader(shaderName+".fs", gl.FRAGMENT_SHADER)
	if err != nil {
		return nil, err
	}

	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return nil, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return (*ShaderProgram)(&program), nil
}

func compileShader(sourcePath string, shaderType uint32) (uint32, error) {
	source, err := shaderResources.ReadFile("shaders/" + sourcePath)
	if err != nil {
		return 0, fmt.Errorf("failed to load shader file %v: %w", sourcePath, err)
	}

	shader := gl.CreateShader(shaderType)
	csources, free := gl.Strs(string(source))
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

		return 0, fmt.Errorf("failed to compile %v: %v", string(source), log)
	}

	return shader, nil
}

func (s *ShaderProgram) Use() {
	gl.UseProgram(uint32(*s))
}

func (s *ShaderProgram) SetUniformMat4(name string, value mgl32.Mat4) {
	gl.UniformMatrix4fv(gl.GetUniformLocation(uint32(*s), gl.Str(name)), 1, false, &value[0])
}

func (s *ShaderProgram) SetUniformInt32(name string, value int32) {
	gl.Uniform1i(gl.GetUniformLocation(uint32(*s), gl.Str(name)), value)
}

func (s *ShaderProgram) SetUniformTexture2D(name string, texture *Texture, n uint32) {
	gl.ActiveTexture(gl.TEXTURE0 + n)
	texture.Bind()
	s.SetUniformInt32(name, int32(n))
}
