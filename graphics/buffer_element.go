package graphics

type BufferElement struct {
	name       string
	sType      ShaderDataType
	size       int32
	offset     uintptr
	normalized bool
}

func NewBufferElement(sType ShaderDataType, name string, normalized bool) *BufferElement {
	return &BufferElement{
		name:       name,
		sType:      sType,
		size:       sType.Size(),
		normalized: normalized,
	}
}

func (element *BufferElement) ComponentCount() int32 {
	return element.sType.ComponentCount()
}
