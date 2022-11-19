package graphics

import (
	"gocraft/gl"
	"unsafe"
)

type IndexBuffer struct {
	rendererID uint32
	count      uint32
}

func NewIndexBuffer(indices []uint32) *IndexBuffer {
	var buffer = &IndexBuffer{count: uint32(len(indices))}
	gl.CreateBuffers(1, &buffer.rendererID)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, buffer.rendererID)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, unsafe.Pointer(&indices[0]), gl.STATIC_DRAW)

	return buffer
}

func (buffer *IndexBuffer) Delete() {
	gl.DeleteBuffers(1, &buffer.rendererID)
}

func (buffer *IndexBuffer) Bind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, buffer.rendererID)
}

func (buffer *IndexBuffer) Unbind() {
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
}

func (buffer *IndexBuffer) Count() uint32 {
	return buffer.count
}
