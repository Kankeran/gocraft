package graphics

import (
	"gocraft/gl"
	"unsafe"
)

type VertexBuffer struct {
	rendererID uint32
	layout     *BufferLayout
}

func NewVertexBuffer(vertices []float32) *VertexBuffer {
	var buffer = &VertexBuffer{}
	gl.CreateBuffers(1, &buffer.rendererID)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer.rendererID)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, unsafe.Pointer(&vertices[0]), gl.STATIC_DRAW)

	return buffer
}

func (buffer *VertexBuffer) Delete() {
	gl.DeleteBuffers(1, &buffer.rendererID)
}

func (buffer *VertexBuffer) Bind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer.rendererID)
}

func (buffer *VertexBuffer) Unbind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func (buffer *VertexBuffer) SetLayout(layout *BufferLayout) {
	buffer.layout = layout
}

func (buffer *VertexBuffer) Layout() *BufferLayout {
	return buffer.layout
}
