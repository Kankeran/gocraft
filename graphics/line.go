package graphics

import (
	"gocraft/gl"

	"github.com/go-gl/mathgl/mgl32"
)

type Line struct {
	vao, vbo uint32
}

func NewLine() *Line {
	var l = &Line{}
	gl.GenVertexArrays(1, &l.vao)
	gl.GenBuffers(1, &l.vbo)

	return l
}

func (l *Line) Render(shader *ShaderProgram, from, to mgl32.Vec3) {
	vertices := []float32{from[0], from[1], from[2], to[0], to[1], to[2]}
	gl.BindVertexArray(l.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, l.vao)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 3*4, 0)
	gl.EnableVertexAttribArray(0)
	shader.SetUniformMat4("model\x00", mgl32.Ident4())

	gl.DrawArrays(gl.LINE_STRIP, 0, 2)
}
