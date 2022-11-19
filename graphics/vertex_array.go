package graphics

import "gocraft/gl"

type VertexArray struct {
	rendererID    uint32
	vertexBuffers []*VertexBuffer
	indexBuffer   *IndexBuffer
}

func NewVertexArray() *VertexArray {
	var vertexArray = &VertexArray{}
	gl.CreateVertexArrays(1, &vertexArray.rendererID)

	return vertexArray
}

func (vertexArray *VertexArray) Delete() {
	gl.DeleteVertexArrays(1, &vertexArray.rendererID)
}

func (vertexArray *VertexArray) Bind() {
	gl.BindVertexArray(vertexArray.rendererID)
}

func (vertexArray *VertexArray) Unbind() {
	gl.BindVertexArray(0)
}

func (vertexArray *VertexArray) AddVertexBuffer(buffer *VertexBuffer) {
	if len(buffer.Layout().Elements()) < 1 {
		panic("VertexBuffer has no layout!")
	}

	vertexArray.Bind()
	buffer.Bind()

	var layout = buffer.Layout()
	for index, element := range layout.Elements() {
		gl.EnableVertexAttribArray(uint32(index))
		gl.VertexAttribPointerWithOffset(
			uint32(index),
			element.ComponentCount(),
			element.sType.OpenGLType(),
			element.normalized,
			layout.Stride(),
			element.offset,
		)
	}

	vertexArray.vertexBuffers = append(vertexArray.vertexBuffers, buffer)
}

func (vertexArray *VertexArray) SetIndexBuffer(buffer *IndexBuffer) {
	vertexArray.Bind()
	buffer.Bind()

	vertexArray.indexBuffer = buffer
}
